package wbwallet

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestNewClientDefaults(t *testing.T) {
	client := NewClient("", nil)

	if client.BaseURL != "https://api.wbpay.ru" {
		t.Fatalf("unexpected BaseURL: %q", client.BaseURL)
	}
	if client.HTTPClient != http.DefaultClient {
		t.Fatalf("expected default http client")
	}
}

func TestGeneratePayerToken(t *testing.T) {
	var gotBody GeneratePayerTokenRequestBody
	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if r.Method != http.MethodPost {
				t.Fatalf("unexpected method: %s", r.Method)
			}
			if r.URL.Path != "/api/v1/users/tokens" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}
			if got := r.Header.Get("Authorization"); got != "Bearer token" {
				t.Fatalf("unexpected Authorization: %q", got)
			}
			if got := r.Header.Get("Content-Type"); got != "application/json" {
				t.Fatalf("unexpected Content-Type: %q", got)
			}
			if got := r.Header.Get("X-Request-Country"); got != "ru" {
				t.Fatalf("unexpected X-Request-Country: %q", got)
			}
			if got := r.Header.Get("X-Request-Region"); got != "ru-mow" {
				t.Fatalf("unexpected X-Request-Region: %q", got)
			}
			if got := r.Header.Get("X-Signature"); got != "signature" {
				t.Fatalf("unexpected X-Signature: %q", got)
			}

			if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
				t.Fatalf("decode request body: %v", err)
			}

			payload, err := json.Marshal(GeneratePayerTokenOKResponse{
				Data: GeneratePayerTokenTokenGenerationData{
					RegistrationID: "reg-1",
					DeepLink:       "wildberries://wbpay/agreement?id=1",
				},
				ErrorCode:        "ERR_NONE",
				ErrorDescription: "ok",
			})
			if err != nil {
				t.Fatalf("marshal response: %v", err)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"X-Test-Header": []string{"ok"},
				},
				Body: io.NopCloser(strings.NewReader(string(payload))),
			}, nil
		}),
	}

	client := NewClient("https://api.wbpay.ru", httpClient)
	result, err := client.GeneratePayerToken(context.Background(), &GeneratePayerTokenRequest{
		Authorization:   "Bearer token",
		ContentType:     "application/json",
		XRequestCountry: "ru",
		XRequestRegion:  "ru-mow",
		XSignature:      "signature",
		Body: GeneratePayerTokenRequestBody{
			TerminalID:  "terminal-1",
			PhoneNumber: "79991234567",
			CreatedAt:   1752057656,
			ClientID:    "client-1",
		},
	})
	if err != nil {
		t.Fatalf("GeneratePayerToken returned error: %v", err)
	}

	if gotBody.TerminalID != "terminal-1" || gotBody.PhoneNumber != "79991234567" || gotBody.CreatedAt != 1752057656 || gotBody.ClientID != "client-1" {
		t.Fatalf("unexpected request body: %+v", gotBody)
	}
	if result.OK == nil {
		t.Fatalf("expected OK response")
	}
	if result.BadRequest != nil || result.Forbidden != nil {
		t.Fatalf("unexpected non-OK response sections: %+v", result)
	}
	if result.OK.Data.RegistrationID != "reg-1" {
		t.Fatalf("unexpected registration id: %q", result.OK.Data.RegistrationID)
	}
	if result.Headers.Get("X-Test-Header") != "ok" {
		t.Fatalf("expected response headers to be preserved")
	}
}

func TestClientAutoFillsAuthorizationAndSignature(t *testing.T) {
	_, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatalf("generate ed25519 key: %v", err)
	}

	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if got := r.Header.Get("Authorization"); got != "Bearer token-value" {
				t.Fatalf("unexpected Authorization: %q", got)
			}
			if got := r.Header.Get("Content-Type"); got != "application/json" {
				t.Fatalf("unexpected Content-Type: %q", got)
			}
			if got := r.Header.Get("X-Request-Country"); got != "ru" {
				t.Fatalf("unexpected X-Request-Country: %q", got)
			}
			if got := r.Header.Get("X-Request-Region"); got != "ru-mow" {
				t.Fatalf("unexpected X-Request-Region: %q", got)
			}

			signature := r.Header.Get("X-Signature")
			if signature == "" {
				t.Fatalf("expected X-Signature to be set")
			}
			if _, err := base64.StdEncoding.DecodeString(signature); err != nil {
				t.Fatalf("signature is not base64: %v", err)
			}

			payload, err := json.Marshal(GeneratePayerTokenOKResponse{
				Data: GeneratePayerTokenTokenGenerationData{
					RegistrationID: "reg-42",
					DeepLink:       "wildberries://wbpay/agreement?id=42",
				},
				ErrorCode:        "ERR_NONE",
				ErrorDescription: "ok",
			})
			if err != nil {
				t.Fatalf("marshal response: %v", err)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader(string(payload))),
			}, nil
		}),
	}

	client := NewClientWithCredentials("https://api.wbpay.ru", httpClient, "token-value", privateKey, "ru", "ru-mow")
	result, err := client.GeneratePayerToken(context.Background(), &GeneratePayerTokenRequest{
		Body: GeneratePayerTokenRequestBody{
			TerminalID:  "terminal-1",
			PhoneNumber: "79991234567",
			CreatedAt:   1752057656,
			ClientID:    "client-42",
		},
	})
	if err != nil {
		t.Fatalf("GeneratePayerToken returned error: %v", err)
	}
	if result.OK == nil {
		t.Fatalf("expected OK response")
	}
	if result.OK.Data.RegistrationID != "reg-42" {
		t.Fatalf("unexpected registration id: %q", result.OK.Data.RegistrationID)
	}
}

func TestClientAutoFillsWBPayID(t *testing.T) {
	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if got := r.Header.Get("Authorization"); got != "Bearer token-value" {
				t.Fatalf("unexpected Authorization: %q", got)
			}
			if got := r.Header.Get("Content-Type"); got != "application/json" {
				t.Fatalf("unexpected Content-Type: %q", got)
			}
			if got := r.Header.Get("X-Wbpay-Id"); got != "order-42" {
				t.Fatalf("unexpected X-Wbpay-Id: %q", got)
			}

			payload, err := json.Marshal(DoOnlinePaymentByPhoneOKResponse{
				Data: DoOnlinePaymentByPhoneDoOrderData{
					DeepLink: "wildberries://wbpay/agreement?id=42",
				},
				ErrorCode:        "ERR_NONE",
				ErrorDescription: "ok",
			})
			if err != nil {
				t.Fatalf("marshal response: %v", err)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader(string(payload))),
			}, nil
		}),
	}

	client := NewClientWithCredentials("https://api.wbpay.ru", httpClient, "token-value", nil, "", "")
	result, err := client.DoOnlinePaymentByPhone(context.Background(), &DoOnlinePaymentByPhoneRequest{
		Body: DoOnlinePaymentByPhoneRequestBody{
			OrderID: "order-42",
		},
	})
	if err != nil {
		t.Fatalf("DoOnlinePaymentByPhone returned error: %v", err)
	}
	if result.OK == nil {
		t.Fatalf("expected OK response")
	}
	if result.OK.Data.DeepLink == "" {
		t.Fatalf("expected deep link in OK response")
	}
}

func TestGetPayerTokenGenerationStatus(t *testing.T) {
	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if r.Method != http.MethodGet {
				t.Fatalf("unexpected method: %s", r.Method)
			}
			if r.URL.EscapedPath() != "/api/v1/users/tokens/id%2Fwith%20spaces/status" {
				t.Fatalf("unexpected path: %s", r.URL.EscapedPath())
			}
			if got := r.Header.Get("X-Wbpay-Id"); got != "id/with spaces" {
				t.Fatalf("unexpected X-Wbpay-Id: %q", got)
			}

			payload, err := json.Marshal(GetPayerTokenGenerationStatusOKResponse{
				Data: GetPayerTokenGenerationStatusTokenStatusData{
					Status:                "succeeded",
					Token:                 "payment-token",
					FailReasonCode:        "USER_NOT_APPROVE",
					FailReasonDescription: "not used in success but documented",
				},
				ErrorCode:        "ERR_NONE",
				ErrorDescription: "ok",
			})
			if err != nil {
				t.Fatalf("marshal response: %v", err)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader(string(payload))),
			}, nil
		}),
	}

	client := NewClient("https://api.wbpay.ru", httpClient)
	result, err := client.GetPayerTokenGenerationStatus(context.Background(), &GetPayerTokenGenerationStatusRequest{
		RegistrationID: "id/with spaces",
		Authorization:  "Bearer token",
		ContentType:    "application/json",
		XWBPayID:       "id/with spaces",
	})
	if err != nil {
		t.Fatalf("GetPayerTokenGenerationStatus returned error: %v", err)
	}
	if result.OK == nil {
		t.Fatalf("expected OK response")
	}
	if result.OK.Data.Status != "succeeded" {
		t.Fatalf("unexpected status: %q", result.OK.Data.Status)
	}
	if result.OK.Data.Token != "payment-token" {
		t.Fatalf("unexpected token: %q", result.OK.Data.Token)
	}
}

func TestDoOfflinePaymentBadRequest(t *testing.T) {
	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			payload, err := json.Marshal(DoOfflinePaymentBadRequestResponse{
				ErrorCode:        "NOT_FOUND",
				ErrorDescription: "order not found",
			})
			if err != nil {
				t.Fatalf("marshal response: %v", err)
			}
			return &http.Response{
				StatusCode: http.StatusBadRequest,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader(string(payload))),
			}, nil
		}),
	}

	client := NewClient("https://api.wbpay.ru", httpClient)
	result, err := client.DoOfflinePayment(context.Background(), &DoOfflinePaymentRequest{
		Authorization: "Bearer token",
		ContentType:   "application/json",
		XWBPayID:      "order-1",
		Body: DoOfflinePaymentRequestBody{
			OrderID: "order-1",
		},
	})
	if err != nil {
		t.Fatalf("DoOfflinePayment returned error: %v", err)
	}
	if result.BadRequest == nil {
		t.Fatalf("expected bad request response")
	}
	if result.BadRequest.ErrorCode != "NOT_FOUND" {
		t.Fatalf("unexpected error code: %q", result.BadRequest.ErrorCode)
	}
	if result.OK != nil || result.Forbidden != nil {
		t.Fatalf("unexpected response sections: %+v", result)
	}
}

func TestCalculateLoyaltyCashbackForbidden(t *testing.T) {
	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusForbidden,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		}),
	}

	client := NewClient("https://api.wbpay.ru", httpClient)
	result, err := client.CalculateLoyaltyCashback(context.Background(), &CalculateLoyaltyCashbackRequest{
		Authorization:   "Bearer token",
		ContentType:     "application/json",
		XRequestCountry: "ru",
		XRequestRegion:  "ru-mow",
		Body: CalculateLoyaltyCashbackRequestBody{
			Amount:       100,
			CurrencyCode: 643,
			TerminalID:   "terminal-1",
		},
	})
	if err != nil {
		t.Fatalf("CalculateLoyaltyCashback returned error: %v", err)
	}
	if result.Forbidden == nil {
		t.Fatalf("expected forbidden response")
	}
	if result.OK != nil || result.BadRequest != nil {
		t.Fatalf("unexpected response sections: %+v", result)
	}
}
