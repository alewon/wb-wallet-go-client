package wbwallet

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var ErrNilRequest = errors.New("wbwallet: nil request")

type UnexpectedStatusError struct {
	StatusCode int
	Body       []byte
}

func (e *UnexpectedStatusError) Error() string {
	return fmt.Sprintf("wbwallet: unexpected status %d: %s", e.StatusCode, string(e.Body))
}

type Client struct {
	BaseURL               string
	HTTPClient            *http.Client
	AuthorizationToken    string
	DefaultContentType    string
	DefaultRequestCountry string
	DefaultRequestRegion  string
	PrivateKey            ed25519.PrivateKey
}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	if baseURL == "" {
		baseURL = "https://api.wbpay.ru"
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		BaseURL:            strings.TrimRight(baseURL, "/"),
		HTTPClient:         httpClient,
		DefaultContentType: "application/json",
	}
}

func NewClientWithCredentials(baseURL string, httpClient *http.Client, authorizationToken string, privateKey ed25519.PrivateKey, requestCountry string, requestRegion string) *Client {
	client := NewClient(baseURL, httpClient)
	client.AuthorizationToken = authorizationToken
	client.PrivateKey = privateKey
	client.DefaultRequestCountry = requestCountry
	client.DefaultRequestRegion = requestRegion
	return client
}

func LoadEd25519PrivateKeyFromPEM(pemData []byte) (ed25519.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("wbwallet: invalid PEM data")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("wbwallet: PEM does not contain an Ed25519 private key")
	}
	return privateKey, nil
}

func LoadEd25519PrivateKeyFromPEMFile(path string) (ed25519.PrivateKey, error) {
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadEd25519PrivateKeyFromPEM(pemData)
}

func (c *Client) authorizationHeader() string {
	if c.AuthorizationToken == "" {
		return ""
	}
	if strings.HasPrefix(c.AuthorizationToken, "Bearer ") {
		return c.AuthorizationToken
	}
	return "Bearer " + c.AuthorizationToken
}

func (c *Client) resolveBodySignature(body any) (string, error) {
	if c.PrivateKey == nil || body == nil {
		return "", nil
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	signature := ed25519.Sign(c.PrivateKey, payload)
	return base64.StdEncoding.EncodeToString(signature), nil
}

func (c *Client) fillCommonHeaders(authorization string, contentType string) (string, string) {
	if authorization == "" {
		authorization = c.authorizationHeader()
	}
	if contentType == "" {
		contentType = c.DefaultContentType
	}
	return authorization, contentType
}

func (c *Client) fillRegionalHeaders(country string, region string) (string, string) {
	if country == "" {
		country = c.DefaultRequestCountry
	}
	if region == "" {
		region = c.DefaultRequestRegion
	}
	return country, region
}

func (c *Client) fillWBPayID(current string, fallback string) string {
	if current != "" {
		return current
	}
	return fallback
}

func (c *Client) fillSignature(signature string, body any) (string, error) {
	if signature != "" {
		return signature, nil
	}
	return c.resolveBodySignature(body)
}

func (c *Client) newRequest(ctx context.Context, method string, path string, body any) (*http.Request, error) {
	var reader io.Reader
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(payload)
	}
	httpRequest, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, reader)
	if err != nil {
		return nil, err
	}
	return httpRequest, nil
}

func (c *Client) do(httpRequest *http.Request) (int, http.Header, []byte, error) {
	response, err := c.HTTPClient.Do(httpRequest)
	if err != nil {
		return 0, nil, nil, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, nil, nil, err
	}
	return response.StatusCode, response.Header.Clone(), responseBody, nil
}

func (c *Client) GeneratePayerToken(ctx context.Context, request *GeneratePayerTokenRequest) (*GeneratePayerTokenResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	country, region := c.fillRegionalHeaders(request.XRequestCountry, request.XRequestRegion)
	signature, err := c.fillSignature(request.XSignature, request.Body)
	if err != nil {
		return nil, err
	}
	path := "/api/v1/users/tokens"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if country != "" {
		httpRequest.Header.Set("X-Request-Country", country)
	}
	if region != "" {
		httpRequest.Header.Set("X-Request-Region", region)
	}
	if signature != "" {
		httpRequest.Header.Set("X-Signature", signature)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GeneratePayerTokenResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GeneratePayerTokenOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GeneratePayerTokenBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GeneratePayerTokenForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetPayerTokenGenerationStatus(ctx context.Context, request *GetPayerTokenGenerationStatusRequest) (*GetPayerTokenGenerationStatusResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.RegistrationID)
	path := "/api/v1/users/tokens/" + url.PathEscape(request.RegistrationID) + "/status"
	httpRequest, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetPayerTokenGenerationStatusResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetPayerTokenGenerationStatusOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetPayerTokenGenerationStatusBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetPayerTokenGenerationStatusForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) RegisterOnlinePaymentByToken(ctx context.Context, request *RegisterOnlinePaymentByTokenRequest) (*RegisterOnlinePaymentByTokenResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	country, region := c.fillRegionalHeaders(request.XRequestCountry, request.XRequestRegion)
	path := "/api/v1/orders/online/register"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if country != "" {
		httpRequest.Header.Set("X-Request-Country", country)
	}
	if region != "" {
		httpRequest.Header.Set("X-Request-Region", region)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &RegisterOnlinePaymentByTokenResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &RegisterOnlinePaymentByTokenOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &RegisterOnlinePaymentByTokenBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &RegisterOnlinePaymentByTokenForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) DoOnlinePaymentByToken(ctx context.Context, request *DoOnlinePaymentByTokenRequest) (*DoOnlinePaymentByTokenResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.Body.OrderID)
	path := "/api/v1/orders/do"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &DoOnlinePaymentByTokenResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &DoOnlinePaymentByTokenOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &DoOnlinePaymentByTokenBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &DoOnlinePaymentByTokenForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetOnlinePaymentByTokenStatus(ctx context.Context, request *GetOnlinePaymentByTokenStatusRequest) (*GetOnlinePaymentByTokenStatusResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.OrderID)
	path := "/api/v1/orders/" + url.PathEscape(request.OrderID) + "/status"
	httpRequest, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetOnlinePaymentByTokenStatusResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetOnlinePaymentByTokenStatusOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetOnlinePaymentByTokenStatusBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetOnlinePaymentByTokenStatusForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) RegisterOnlinePaymentByPhone(ctx context.Context, request *RegisterOnlinePaymentByPhoneRequest) (*RegisterOnlinePaymentByPhoneResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	country, region := c.fillRegionalHeaders(request.XRequestCountry, request.XRequestRegion)
	signature, err := c.fillSignature(request.XSignature, request.Body)
	if err != nil {
		return nil, err
	}
	path := "/api/v1/orders/online/register-by-phone"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if country != "" {
		httpRequest.Header.Set("X-Request-Country", country)
	}
	if region != "" {
		httpRequest.Header.Set("X-Request-Region", region)
	}
	if signature != "" {
		httpRequest.Header.Set("X-Signature", signature)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &RegisterOnlinePaymentByPhoneResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &RegisterOnlinePaymentByPhoneOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &RegisterOnlinePaymentByPhoneBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &RegisterOnlinePaymentByPhoneForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) DoOnlinePaymentByPhone(ctx context.Context, request *DoOnlinePaymentByPhoneRequest) (*DoOnlinePaymentByPhoneResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.Body.OrderID)
	path := "/api/v1/orders/do"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &DoOnlinePaymentByPhoneResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &DoOnlinePaymentByPhoneOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &DoOnlinePaymentByPhoneBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &DoOnlinePaymentByPhoneForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetOnlinePaymentByPhoneStatus(ctx context.Context, request *GetOnlinePaymentByPhoneStatusRequest) (*GetOnlinePaymentByPhoneStatusResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.OrderID)
	path := "/api/v1/orders/" + url.PathEscape(request.OrderID) + "/status"
	httpRequest, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetOnlinePaymentByPhoneStatusResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetOnlinePaymentByPhoneStatusOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetOnlinePaymentByPhoneStatusBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetOnlinePaymentByPhoneStatusForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) RegisterPaymentLink(ctx context.Context, request *RegisterPaymentLinkRequest) (*RegisterPaymentLinkResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	country, region := c.fillRegionalHeaders(request.XRequestCountry, request.XRequestRegion)
	signature, err := c.fillSignature(request.XSignature, request.Body)
	if err != nil {
		return nil, err
	}
	path := "/api/v1/orders/hpp/register"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if country != "" {
		httpRequest.Header.Set("X-Request-Country", country)
	}
	if region != "" {
		httpRequest.Header.Set("X-Request-Region", region)
	}
	if signature != "" {
		httpRequest.Header.Set("X-Signature", signature)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &RegisterPaymentLinkResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &RegisterPaymentLinkOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &RegisterPaymentLinkBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &RegisterPaymentLinkForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetPaymentLinkStatus(ctx context.Context, request *GetPaymentLinkStatusRequest) (*GetPaymentLinkStatusResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.OrderID)
	path := "/api/v1/orders/" + url.PathEscape(request.OrderID) + "/status"
	httpRequest, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetPaymentLinkStatusResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetPaymentLinkStatusOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetPaymentLinkStatusBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetPaymentLinkStatusForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) RegisterOnlinePaymentWithTokenCreation(ctx context.Context, request *RegisterOnlinePaymentWithTokenCreationRequest) (*RegisterOnlinePaymentWithTokenCreationResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	country, region := c.fillRegionalHeaders(request.XRequestCountry, request.XRequestRegion)
	signature, err := c.fillSignature(request.XSignature, request.Body)
	if err != nil {
		return nil, err
	}
	path := "/api/v1/orders/online/register-with-create-token"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if country != "" {
		httpRequest.Header.Set("X-Request-Country", country)
	}
	if region != "" {
		httpRequest.Header.Set("X-Request-Region", region)
	}
	if signature != "" {
		httpRequest.Header.Set("X-Signature", signature)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &RegisterOnlinePaymentWithTokenCreationResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &RegisterOnlinePaymentWithTokenCreationOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &RegisterOnlinePaymentWithTokenCreationBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &RegisterOnlinePaymentWithTokenCreationForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) DoOnlinePaymentWithTokenCreation(ctx context.Context, request *DoOnlinePaymentWithTokenCreationRequest) (*DoOnlinePaymentWithTokenCreationResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.Body.OrderID)
	path := "/api/v1/orders/do"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &DoOnlinePaymentWithTokenCreationResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &DoOnlinePaymentWithTokenCreationOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &DoOnlinePaymentWithTokenCreationBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &DoOnlinePaymentWithTokenCreationForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetOnlinePaymentWithTokenCreationStatus(ctx context.Context, request *GetOnlinePaymentWithTokenCreationStatusRequest) (*GetOnlinePaymentWithTokenCreationStatusResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.OrderID)
	path := "/api/v1/orders/" + url.PathEscape(request.OrderID) + "/status"
	httpRequest, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetOnlinePaymentWithTokenCreationStatusResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetOnlinePaymentWithTokenCreationStatusOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetOnlinePaymentWithTokenCreationStatusBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetOnlinePaymentWithTokenCreationStatusForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) RegisterOnlineRefund(ctx context.Context, request *RegisterOnlineRefundRequest) (*RegisterOnlineRefundResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	signature, err := c.fillSignature(request.XSignature, request.Body)
	if err != nil {
		return nil, err
	}
	wbPayID := c.fillWBPayID(request.XWBPayID, request.Body.OrderID)
	path := "/api/v1/refunds/register"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if signature != "" {
		httpRequest.Header.Set("X-Signature", signature)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &RegisterOnlineRefundResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &RegisterOnlineRefundOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &RegisterOnlineRefundBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &RegisterOnlineRefundForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) DoOnlineRefund(ctx context.Context, request *DoOnlineRefundRequest) (*DoOnlineRefundResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.Body.RefundID)
	path := "/api/v1/refunds/do"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &DoOnlineRefundResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &DoOnlineRefundOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &DoOnlineRefundBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &DoOnlineRefundForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetOnlineRefundStatus(ctx context.Context, request *GetOnlineRefundStatusRequest) (*GetOnlineRefundStatusResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.RefundID)
	path := "/api/v1/refunds/" + url.PathEscape(request.RefundID) + "/status"
	httpRequest, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetOnlineRefundStatusResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetOnlineRefundStatusOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetOnlineRefundStatusBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetOnlineRefundStatusForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) RegisterOfflinePayment(ctx context.Context, request *RegisterOfflinePaymentRequest) (*RegisterOfflinePaymentResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	country, region := c.fillRegionalHeaders(request.XRequestCountry, request.XRequestRegion)
	signature, err := c.fillSignature(request.XSignature, request.Body)
	if err != nil {
		return nil, err
	}
	path := "/api/v1/orders/offline/register"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if country != "" {
		httpRequest.Header.Set("X-Request-Country", country)
	}
	if region != "" {
		httpRequest.Header.Set("X-Request-Region", region)
	}
	if signature != "" {
		httpRequest.Header.Set("X-Signature", signature)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &RegisterOfflinePaymentResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &RegisterOfflinePaymentOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &RegisterOfflinePaymentBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &RegisterOfflinePaymentForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) DoOfflinePayment(ctx context.Context, request *DoOfflinePaymentRequest) (*DoOfflinePaymentResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.Body.OrderID)
	path := "/api/v1/orders/do"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &DoOfflinePaymentResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &DoOfflinePaymentOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &DoOfflinePaymentBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &DoOfflinePaymentForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetOfflinePaymentStatus(ctx context.Context, request *GetOfflinePaymentStatusRequest) (*GetOfflinePaymentStatusResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.OrderID)
	path := "/api/v1/orders/" + url.PathEscape(request.OrderID) + "/status"
	httpRequest, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetOfflinePaymentStatusResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetOfflinePaymentStatusOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetOfflinePaymentStatusBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetOfflinePaymentStatusForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) RegisterOfflineRefund(ctx context.Context, request *RegisterOfflineRefundRequest) (*RegisterOfflineRefundResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	signature, err := c.fillSignature(request.XSignature, request.Body)
	if err != nil {
		return nil, err
	}
	wbPayID := c.fillWBPayID(request.XWBPayID, request.Body.OrderID)
	path := "/api/v1/refunds/register"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if signature != "" {
		httpRequest.Header.Set("X-Signature", signature)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &RegisterOfflineRefundResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &RegisterOfflineRefundOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &RegisterOfflineRefundBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &RegisterOfflineRefundForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) DoOfflineRefund(ctx context.Context, request *DoOfflineRefundRequest) (*DoOfflineRefundResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.Body.RefundID)
	path := "/api/v1/refunds/do"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &DoOfflineRefundResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &DoOfflineRefundOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &DoOfflineRefundBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &DoOfflineRefundForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetOfflineRefundStatus(ctx context.Context, request *GetOfflineRefundStatusRequest) (*GetOfflineRefundStatusResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	wbPayID := c.fillWBPayID(request.XWBPayID, request.RefundID)
	path := "/api/v1/refunds/" + url.PathEscape(request.RefundID) + "/status"
	httpRequest, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if wbPayID != "" {
		httpRequest.Header.Set("X-Wbpay-Id", wbPayID)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetOfflineRefundStatusResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetOfflineRefundStatusOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetOfflineRefundStatusBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetOfflineRefundStatusForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) GetLoyaltyTerminalSettings(ctx context.Context, request *GetLoyaltyTerminalSettingsRequest) (*GetLoyaltyTerminalSettingsResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	country, region := c.fillRegionalHeaders(request.XRequestCountry, request.XRequestRegion)
	path := "/api/v1/loyalties/settings"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if country != "" {
		httpRequest.Header.Set("X-Request-Country", country)
	}
	if region != "" {
		httpRequest.Header.Set("X-Request-Region", region)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &GetLoyaltyTerminalSettingsResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &GetLoyaltyTerminalSettingsOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &GetLoyaltyTerminalSettingsBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &GetLoyaltyTerminalSettingsForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}

func (c *Client) CalculateLoyaltyCashback(ctx context.Context, request *CalculateLoyaltyCashbackRequest) (*CalculateLoyaltyCashbackResult, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	authorization, contentType := c.fillCommonHeaders(request.Authorization, request.ContentType)
	country, region := c.fillRegionalHeaders(request.XRequestCountry, request.XRequestRegion)
	path := "/api/v1/loyalties/calculate_cashback"
	httpRequest, err := c.newRequest(ctx, "POST", path, request.Body)
	if err != nil {
		return nil, err
	}
	if authorization != "" {
		httpRequest.Header.Set("Authorization", authorization)
	}
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	if country != "" {
		httpRequest.Header.Set("X-Request-Country", country)
	}
	if region != "" {
		httpRequest.Header.Set("X-Request-Region", region)
	}
	statusCode, responseHeaders, responseBody, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	result := &CalculateLoyaltyCashbackResult{StatusCode: statusCode, Headers: responseHeaders, RawBody: responseBody}
	switch statusCode {
	case http.StatusOK:
		result.OK = &CalculateLoyaltyCashbackOKResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.OK); err != nil {
				return nil, err
			}
		}
	case http.StatusBadRequest:
		result.BadRequest = &CalculateLoyaltyCashbackBadRequestResponse{}
		if len(responseBody) > 0 {
			if err := json.Unmarshal(responseBody, result.BadRequest); err != nil {
				return nil, err
			}
		}
	case http.StatusForbidden:
		result.Forbidden = &CalculateLoyaltyCashbackForbiddenResponse{}
	default:
		return nil, &UnexpectedStatusError{StatusCode: statusCode, Body: responseBody}
	}
	return result, nil
}
