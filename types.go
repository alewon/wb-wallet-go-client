package wbwallet

import "net/http"

type GeneratePayerTokenRequest struct {
	Authorization   string                        `json:"-"`
	ContentType     string                        `json:"-"`
	XRequestCountry string                        `json:"-"`
	XRequestRegion  string                        `json:"-"`
	XSignature      string                        `json:"-"`
	Body            GeneratePayerTokenRequestBody `json:"-"`
}

type GeneratePayerTokenRequestBody struct {
	ClientID    string `json:"client_id,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	TerminalID  string `json:"terminal_id,omitempty"`
}

type GeneratePayerTokenResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GeneratePayerTokenOKResponse
	BadRequest *GeneratePayerTokenBadRequestResponse
	Forbidden  *GeneratePayerTokenForbiddenResponse
}

type GeneratePayerTokenForbiddenResponse struct{}

type GeneratePayerTokenOKResponse struct {
	Data             GeneratePayerTokenTokenGenerationData `json:"data,omitempty"`
	ErrorCode        string                                `json:"error_code,omitempty"`
	ErrorDescription string                                `json:"error_description,omitempty"`
}

type GeneratePayerTokenBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GeneratePayerTokenTokenGenerationData struct {
	RegistrationID string `json:"registration_id,omitempty"`
	DeepLink       string `json:"deep_link,omitempty"`
}

type GetPayerTokenGenerationStatusRequest struct {
	RegistrationID string `json:"-"`
	Authorization  string `json:"-"`
	ContentType    string `json:"-"`
	XWBPayID       string `json:"-"`
}

type GetPayerTokenGenerationStatusResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetPayerTokenGenerationStatusOKResponse
	BadRequest *GetPayerTokenGenerationStatusBadRequestResponse
	Forbidden  *GetPayerTokenGenerationStatusForbiddenResponse
}

type GetPayerTokenGenerationStatusForbiddenResponse struct{}

type GetPayerTokenGenerationStatusOKResponse struct {
	Data             GetPayerTokenGenerationStatusTokenStatusData `json:"data,omitempty"`
	ErrorCode        string                                       `json:"error_code,omitempty"`
	ErrorDescription string                                       `json:"error_description,omitempty"`
}

type GetPayerTokenGenerationStatusBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetPayerTokenGenerationStatusTokenStatusData struct {
	Status                string `json:"status,omitempty"`
	FailReasonCode        string `json:"fail_reason_code,omitempty"`
	FailReasonDescription string `json:"fail_reason_description,omitempty"`
	Token                 string `json:"token,omitempty"`
}

type RegisterOnlinePaymentByTokenRequest struct {
	Authorization   string                                  `json:"-"`
	ContentType     string                                  `json:"-"`
	XRequestCountry string                                  `json:"-"`
	XRequestRegion  string                                  `json:"-"`
	Body            RegisterOnlinePaymentByTokenRequestBody `json:"-"`
}

type RegisterOnlinePaymentByTokenRequestBody struct {
	Amount       int64                                         `json:"amount,omitempty"`
	ClientID     string                                        `json:"client_id,omitempty"`
	CreatedAt    int64                                         `json:"created_at,omitempty"`
	CurrencyCode int64                                         `json:"currency_code,omitempty"`
	TerminalID   string                                        `json:"terminal_id,omitempty"`
	Token        string                                        `json:"token,omitempty"`
	InvoiceID    string                                        `json:"invoice_id,omitempty"`
	Positions    []RegisterOnlinePaymentByTokenRequestPosition `json:"positions,omitempty"`
	RedirectURL  string                                        `json:"redirect_url,omitempty"`
}

type RegisterOnlinePaymentByTokenRequestPosition struct {
	Count int64  `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type RegisterOnlinePaymentByTokenResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *RegisterOnlinePaymentByTokenOKResponse
	BadRequest *RegisterOnlinePaymentByTokenBadRequestResponse
	Forbidden  *RegisterOnlinePaymentByTokenForbiddenResponse
}

type RegisterOnlinePaymentByTokenForbiddenResponse struct{}

type RegisterOnlinePaymentByTokenOKResponse struct {
	Data             RegisterOnlinePaymentByTokenOrderRegistrationData `json:"data,omitempty"`
	ErrorCode        string                                            `json:"error_code,omitempty"`
	ErrorDescription string                                            `json:"error_description,omitempty"`
}

type RegisterOnlinePaymentByTokenBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type RegisterOnlinePaymentByTokenOrderRegistrationData struct {
	OrderID string `json:"order_id,omitempty"`
}

type DoOnlinePaymentByTokenRequest struct {
	Authorization string                            `json:"-"`
	ContentType   string                            `json:"-"`
	XWBPayID      string                            `json:"-"`
	Body          DoOnlinePaymentByTokenRequestBody `json:"-"`
}

type DoOnlinePaymentByTokenRequestBody struct {
	OrderID string `json:"order_id,omitempty"`
}

type DoOnlinePaymentByTokenResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *DoOnlinePaymentByTokenOKResponse
	BadRequest *DoOnlinePaymentByTokenBadRequestResponse
	Forbidden  *DoOnlinePaymentByTokenForbiddenResponse
}

type DoOnlinePaymentByTokenForbiddenResponse struct{}

type DoOnlinePaymentByTokenOKResponse struct {
	Data             DoOnlinePaymentByTokenDoOrderData `json:"data,omitempty"`
	ErrorCode        string                            `json:"error_code,omitempty"`
	ErrorDescription string                            `json:"error_description,omitempty"`
}

type DoOnlinePaymentByTokenBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type DoOnlinePaymentByTokenDoOrderData struct {
	DeepLink string `json:"deep_link,omitempty"`
}

type GetOnlinePaymentByTokenStatusRequest struct {
	OrderID       string `json:"-"`
	Authorization string `json:"-"`
	ContentType   string `json:"-"`
	XWBPayID      string `json:"-"`
}

type GetOnlinePaymentByTokenStatusResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetOnlinePaymentByTokenStatusOKResponse
	BadRequest *GetOnlinePaymentByTokenStatusBadRequestResponse
	Forbidden  *GetOnlinePaymentByTokenStatusForbiddenResponse
}

type GetOnlinePaymentByTokenStatusForbiddenResponse struct{}

type GetOnlinePaymentByTokenStatusOKResponse struct {
	Data             GetOnlinePaymentByTokenStatusOrderStatusData `json:"data,omitempty"`
	ErrorCode        string                                       `json:"error_code,omitempty"`
	ErrorDescription string                                       `json:"error_description,omitempty"`
}

type GetOnlinePaymentByTokenStatusBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOnlinePaymentByTokenStatusOrderStatusData struct {
	Status                string `json:"status,omitempty"`
	FailReasonCode        string `json:"fail_reason_code,omitempty"`
	FailReasonDescription string `json:"fail_reason_description,omitempty"`
	Token                 string `json:"token,omitempty"`
}

type RegisterOnlinePaymentByPhoneRequest struct {
	Authorization   string                                  `json:"-"`
	ContentType     string                                  `json:"-"`
	XRequestCountry string                                  `json:"-"`
	XRequestRegion  string                                  `json:"-"`
	XSignature      string                                  `json:"-"`
	Body            RegisterOnlinePaymentByPhoneRequestBody `json:"-"`
}

type RegisterOnlinePaymentByPhoneRequestBody struct {
	Amount       int64                                         `json:"amount,omitempty"`
	CreatedAt    int64                                         `json:"created_at,omitempty"`
	CurrencyCode int64                                         `json:"currency_code,omitempty"`
	PhoneNumber  string                                        `json:"phone_number,omitempty"`
	TerminalID   string                                        `json:"terminal_id,omitempty"`
	InvoiceID    string                                        `json:"invoice_id,omitempty"`
	Positions    []RegisterOnlinePaymentByPhoneRequestPosition `json:"positions,omitempty"`
	RedirectURL  string                                        `json:"redirect_url,omitempty"`
}

type RegisterOnlinePaymentByPhoneRequestPosition struct {
	Count int64  `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type RegisterOnlinePaymentByPhoneResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *RegisterOnlinePaymentByPhoneOKResponse
	BadRequest *RegisterOnlinePaymentByPhoneBadRequestResponse
	Forbidden  *RegisterOnlinePaymentByPhoneForbiddenResponse
}

type RegisterOnlinePaymentByPhoneForbiddenResponse struct{}

type RegisterOnlinePaymentByPhoneOKResponse struct {
	Data             RegisterOnlinePaymentByPhoneOrderRegistrationData `json:"data,omitempty"`
	ErrorCode        string                                            `json:"error_code,omitempty"`
	ErrorDescription string                                            `json:"error_description,omitempty"`
}

type RegisterOnlinePaymentByPhoneBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type RegisterOnlinePaymentByPhoneOrderRegistrationData struct {
	OrderID string `json:"order_id,omitempty"`
}

type DoOnlinePaymentByPhoneRequest struct {
	Authorization string                            `json:"-"`
	ContentType   string                            `json:"-"`
	XWBPayID      string                            `json:"-"`
	Body          DoOnlinePaymentByPhoneRequestBody `json:"-"`
}

type DoOnlinePaymentByPhoneRequestBody struct {
	OrderID string `json:"order_id,omitempty"`
}

type DoOnlinePaymentByPhoneResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *DoOnlinePaymentByPhoneOKResponse
	BadRequest *DoOnlinePaymentByPhoneBadRequestResponse
	Forbidden  *DoOnlinePaymentByPhoneForbiddenResponse
}

type DoOnlinePaymentByPhoneForbiddenResponse struct{}

type DoOnlinePaymentByPhoneOKResponse struct {
	Data             DoOnlinePaymentByPhoneDoOrderData `json:"data,omitempty"`
	ErrorCode        string                            `json:"error_code,omitempty"`
	ErrorDescription string                            `json:"error_description,omitempty"`
}

type DoOnlinePaymentByPhoneBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type DoOnlinePaymentByPhoneDoOrderData struct {
	DeepLink string `json:"deep_link,omitempty"`
}

type GetOnlinePaymentByPhoneStatusRequest struct {
	OrderID       string `json:"-"`
	Authorization string `json:"-"`
	ContentType   string `json:"-"`
	XWBPayID      string `json:"-"`
}

type GetOnlinePaymentByPhoneStatusResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetOnlinePaymentByPhoneStatusOKResponse
	BadRequest *GetOnlinePaymentByPhoneStatusBadRequestResponse
	Forbidden  *GetOnlinePaymentByPhoneStatusForbiddenResponse
}

type GetOnlinePaymentByPhoneStatusForbiddenResponse struct{}

type GetOnlinePaymentByPhoneStatusOKResponse struct {
	Data             GetOnlinePaymentByPhoneStatusOrderStatusData `json:"data,omitempty"`
	ErrorCode        string                                       `json:"error_code,omitempty"`
	ErrorDescription string                                       `json:"error_description,omitempty"`
}

type GetOnlinePaymentByPhoneStatusBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOnlinePaymentByPhoneStatusOrderStatusData struct {
	Status                string `json:"status,omitempty"`
	FailReasonCode        string `json:"fail_reason_code,omitempty"`
	FailReasonDescription string `json:"fail_reason_description,omitempty"`
	Token                 string `json:"token,omitempty"`
}

type RegisterPaymentLinkRequest struct {
	Authorization   string                         `json:"-"`
	ContentType     string                         `json:"-"`
	XRequestCountry string                         `json:"-"`
	XRequestRegion  string                         `json:"-"`
	XSignature      string                         `json:"-"`
	Body            RegisterPaymentLinkRequestBody `json:"-"`
}

type RegisterPaymentLinkRequestBody struct {
	Amount       int64                                `json:"amount,omitempty"`
	CreatedAt    int64                                `json:"created_at,omitempty"`
	CurrencyCode int64                                `json:"currency_code,omitempty"`
	PhoneNumber  string                               `json:"phone_number,omitempty"`
	RedirectURL  string                               `json:"redirect_url,omitempty"`
	TerminalID   string                               `json:"terminal_id,omitempty"`
	InvoiceID    string                               `json:"invoice_id,omitempty"`
	Positions    []RegisterPaymentLinkRequestPosition `json:"positions,omitempty"`
}

type RegisterPaymentLinkRequestPosition struct {
	Count int64  `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type RegisterPaymentLinkResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *RegisterPaymentLinkOKResponse
	BadRequest *RegisterPaymentLinkBadRequestResponse
	Forbidden  *RegisterPaymentLinkForbiddenResponse
}

type RegisterPaymentLinkForbiddenResponse struct{}

type RegisterPaymentLinkOKResponse struct {
	Data             RegisterPaymentLinkHPPOrderRegistrationData `json:"data,omitempty"`
	ErrorCode        string                                      `json:"error_code,omitempty"`
	ErrorDescription string                                      `json:"error_description,omitempty"`
}

type RegisterPaymentLinkBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type RegisterPaymentLinkHPPOrderRegistrationData struct {
	OrderID    string `json:"order_id,omitempty"`
	PaymentURL string `json:"payment_url,omitempty"`
}

type GetPaymentLinkStatusRequest struct {
	OrderID       string `json:"-"`
	Authorization string `json:"-"`
	ContentType   string `json:"-"`
	XWBPayID      string `json:"-"`
}

type GetPaymentLinkStatusResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetPaymentLinkStatusOKResponse
	BadRequest *GetPaymentLinkStatusBadRequestResponse
	Forbidden  *GetPaymentLinkStatusForbiddenResponse
}

type GetPaymentLinkStatusForbiddenResponse struct{}

type GetPaymentLinkStatusOKResponse struct {
	Data             GetPaymentLinkStatusOrderStatusData `json:"data,omitempty"`
	ErrorCode        string                              `json:"error_code,omitempty"`
	ErrorDescription string                              `json:"error_description,omitempty"`
}

type GetPaymentLinkStatusBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetPaymentLinkStatusOrderStatusData struct {
	Status                string `json:"status,omitempty"`
	FailReasonCode        string `json:"fail_reason_code,omitempty"`
	FailReasonDescription string `json:"fail_reason_description,omitempty"`
	Token                 string `json:"token,omitempty"`
}

type RegisterOnlinePaymentWithTokenCreationRequest struct {
	Authorization   string                                            `json:"-"`
	ContentType     string                                            `json:"-"`
	XRequestCountry string                                            `json:"-"`
	XRequestRegion  string                                            `json:"-"`
	XSignature      string                                            `json:"-"`
	Body            RegisterOnlinePaymentWithTokenCreationRequestBody `json:"-"`
}

type RegisterOnlinePaymentWithTokenCreationRequestBody struct {
	Amount       int64                                                   `json:"amount,omitempty"`
	ClientID     string                                                  `json:"client_id,omitempty"`
	CreatedAt    int64                                                   `json:"created_at,omitempty"`
	CurrencyCode int64                                                   `json:"currency_code,omitempty"`
	PhoneNumber  string                                                  `json:"phone_number,omitempty"`
	TerminalID   string                                                  `json:"terminal_id,omitempty"`
	InvoiceID    string                                                  `json:"invoice_id,omitempty"`
	Positions    []RegisterOnlinePaymentWithTokenCreationRequestPosition `json:"positions,omitempty"`
	RedirectURL  string                                                  `json:"redirect_url,omitempty"`
}

type RegisterOnlinePaymentWithTokenCreationRequestPosition struct {
	Count int64  `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type RegisterOnlinePaymentWithTokenCreationResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *RegisterOnlinePaymentWithTokenCreationOKResponse
	BadRequest *RegisterOnlinePaymentWithTokenCreationBadRequestResponse
	Forbidden  *RegisterOnlinePaymentWithTokenCreationForbiddenResponse
}

type RegisterOnlinePaymentWithTokenCreationForbiddenResponse struct{}

type RegisterOnlinePaymentWithTokenCreationOKResponse struct {
	Data             RegisterOnlinePaymentWithTokenCreationOrderRegistrationData `json:"data,omitempty"`
	ErrorCode        string                                                      `json:"error_code,omitempty"`
	ErrorDescription string                                                      `json:"error_description,omitempty"`
}

type RegisterOnlinePaymentWithTokenCreationBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type RegisterOnlinePaymentWithTokenCreationOrderRegistrationData struct {
	OrderID string `json:"order_id,omitempty"`
}

type DoOnlinePaymentWithTokenCreationRequest struct {
	Authorization string                                      `json:"-"`
	ContentType   string                                      `json:"-"`
	XWBPayID      string                                      `json:"-"`
	Body          DoOnlinePaymentWithTokenCreationRequestBody `json:"-"`
}

type DoOnlinePaymentWithTokenCreationRequestBody struct {
	OrderID string `json:"order_id,omitempty"`
}

type DoOnlinePaymentWithTokenCreationResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *DoOnlinePaymentWithTokenCreationOKResponse
	BadRequest *DoOnlinePaymentWithTokenCreationBadRequestResponse
	Forbidden  *DoOnlinePaymentWithTokenCreationForbiddenResponse
}

type DoOnlinePaymentWithTokenCreationForbiddenResponse struct{}

type DoOnlinePaymentWithTokenCreationOKResponse struct {
	Data             DoOnlinePaymentWithTokenCreationDoOrderData `json:"data,omitempty"`
	ErrorCode        string                                      `json:"error_code,omitempty"`
	ErrorDescription string                                      `json:"error_description,omitempty"`
}

type DoOnlinePaymentWithTokenCreationBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type DoOnlinePaymentWithTokenCreationDoOrderData struct {
	DeepLink string `json:"deep_link,omitempty"`
}

type GetOnlinePaymentWithTokenCreationStatusRequest struct {
	OrderID       string `json:"-"`
	Authorization string `json:"-"`
	ContentType   string `json:"-"`
	XWBPayID      string `json:"-"`
}

type GetOnlinePaymentWithTokenCreationStatusResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetOnlinePaymentWithTokenCreationStatusOKResponse
	BadRequest *GetOnlinePaymentWithTokenCreationStatusBadRequestResponse
	Forbidden  *GetOnlinePaymentWithTokenCreationStatusForbiddenResponse
}

type GetOnlinePaymentWithTokenCreationStatusForbiddenResponse struct{}

type GetOnlinePaymentWithTokenCreationStatusOKResponse struct {
	Data             GetOnlinePaymentWithTokenCreationStatusOrderStatusData `json:"data,omitempty"`
	ErrorCode        string                                                 `json:"error_code,omitempty"`
	ErrorDescription string                                                 `json:"error_description,omitempty"`
}

type GetOnlinePaymentWithTokenCreationStatusBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOnlinePaymentWithTokenCreationStatusOrderStatusData struct {
	Status                string `json:"status,omitempty"`
	FailReasonCode        string `json:"fail_reason_code,omitempty"`
	FailReasonDescription string `json:"fail_reason_description,omitempty"`
	Token                 string `json:"token,omitempty"`
}

type RegisterOnlineRefundRequest struct {
	Authorization string                          `json:"-"`
	ContentType   string                          `json:"-"`
	XSignature    string                          `json:"-"`
	XWBPayID      string                          `json:"-"`
	Body          RegisterOnlineRefundRequestBody `json:"-"`
}

type RegisterOnlineRefundRequestBody struct {
	Amount       int64                                 `json:"amount,omitempty"`
	CreatedAt    int64                                 `json:"created_at,omitempty"`
	CurrencyCode int64                                 `json:"currency_code,omitempty"`
	OrderID      string                                `json:"order_id,omitempty"`
	TerminalID   string                                `json:"terminal_id,omitempty"`
	InvoiceID    string                                `json:"invoice_id,omitempty"`
	Positions    []RegisterOnlineRefundRequestPosition `json:"positions,omitempty"`
}

type RegisterOnlineRefundRequestPosition struct {
	Count int64  `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type RegisterOnlineRefundResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *RegisterOnlineRefundOKResponse
	BadRequest *RegisterOnlineRefundBadRequestResponse
	Forbidden  *RegisterOnlineRefundForbiddenResponse
}

type RegisterOnlineRefundForbiddenResponse struct{}

type RegisterOnlineRefundOKResponse struct {
	Data             RegisterOnlineRefundRegisterRefundData `json:"data,omitempty"`
	ErrorCode        string                                 `json:"error_code,omitempty"`
	ErrorDescription string                                 `json:"error_description,omitempty"`
}

type RegisterOnlineRefundBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type RegisterOnlineRefundRegisterRefundData struct {
	RefundID string `json:"refund_id,omitempty"`
}

type DoOnlineRefundRequest struct {
	Authorization string                    `json:"-"`
	ContentType   string                    `json:"-"`
	XWBPayID      string                    `json:"-"`
	Body          DoOnlineRefundRequestBody `json:"-"`
}

type DoOnlineRefundRequestBody struct {
	RefundID string `json:"refund_id,omitempty"`
}

type DoOnlineRefundResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *DoOnlineRefundOKResponse
	BadRequest *DoOnlineRefundBadRequestResponse
	Forbidden  *DoOnlineRefundForbiddenResponse
}

type DoOnlineRefundForbiddenResponse struct{}

type DoOnlineRefundOKResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type DoOnlineRefundBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOnlineRefundStatusRequest struct {
	RefundID      string `json:"-"`
	Authorization string `json:"-"`
	ContentType   string `json:"-"`
	XWBPayID      string `json:"-"`
}

type GetOnlineRefundStatusResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetOnlineRefundStatusOKResponse
	BadRequest *GetOnlineRefundStatusBadRequestResponse
	Forbidden  *GetOnlineRefundStatusForbiddenResponse
}

type GetOnlineRefundStatusForbiddenResponse struct{}

type GetOnlineRefundStatusOKResponse struct {
	Data             GetOnlineRefundStatusRefundStatusData `json:"data,omitempty"`
	ErrorCode        string                                `json:"error_code,omitempty"`
	ErrorDescription string                                `json:"error_description,omitempty"`
}

type GetOnlineRefundStatusBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOnlineRefundStatusRefundStatusData struct {
	Status                string `json:"status,omitempty"`
	FailReasonCode        string `json:"fail_reason_code,omitempty"`
	FailReasonDescription string `json:"fail_reason_description,omitempty"`
}

type RegisterOfflinePaymentRequest struct {
	Authorization   string                            `json:"-"`
	ContentType     string                            `json:"-"`
	XRequestCountry string                            `json:"-"`
	XRequestRegion  string                            `json:"-"`
	XSignature      string                            `json:"-"`
	Body            RegisterOfflinePaymentRequestBody `json:"-"`
}

type RegisterOfflinePaymentRequestBody struct {
	Amount       int64                                   `json:"amount,omitempty"`
	CreatedAt    int64                                   `json:"created_at,omitempty"`
	CurrencyCode int64                                   `json:"currency_code,omitempty"`
	QRCode       string                                  `json:"qr_code,omitempty"`
	TerminalID   string                                  `json:"terminal_id,omitempty"`
	InvoiceID    string                                  `json:"invoice_id,omitempty"`
	Positions    []RegisterOfflinePaymentRequestPosition `json:"positions,omitempty"`
}

type RegisterOfflinePaymentRequestPosition struct {
	Count int64  `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type RegisterOfflinePaymentResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *RegisterOfflinePaymentOKResponse
	BadRequest *RegisterOfflinePaymentBadRequestResponse
	Forbidden  *RegisterOfflinePaymentForbiddenResponse
}

type RegisterOfflinePaymentForbiddenResponse struct{}

type RegisterOfflinePaymentOKResponse struct {
	Data             RegisterOfflinePaymentRegisterOfflineOrderData `json:"data,omitempty"`
	ErrorCode        string                                         `json:"error_code,omitempty"`
	ErrorDescription string                                         `json:"error_description,omitempty"`
}

type RegisterOfflinePaymentBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type RegisterOfflinePaymentRegisterOfflineOrderData struct {
	OrderID string `json:"order_id,omitempty"`
}

type DoOfflinePaymentRequest struct {
	Authorization string                      `json:"-"`
	ContentType   string                      `json:"-"`
	XWBPayID      string                      `json:"-"`
	Body          DoOfflinePaymentRequestBody `json:"-"`
}

type DoOfflinePaymentRequestBody struct {
	OrderID string `json:"order_id,omitempty"`
}

type DoOfflinePaymentResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *DoOfflinePaymentOKResponse
	BadRequest *DoOfflinePaymentBadRequestResponse
	Forbidden  *DoOfflinePaymentForbiddenResponse
}

type DoOfflinePaymentForbiddenResponse struct{}

type DoOfflinePaymentOKResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type DoOfflinePaymentBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOfflinePaymentStatusRequest struct {
	OrderID       string `json:"-"`
	Authorization string `json:"-"`
	ContentType   string `json:"-"`
	XWBPayID      string `json:"-"`
}

type GetOfflinePaymentStatusResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetOfflinePaymentStatusOKResponse
	BadRequest *GetOfflinePaymentStatusBadRequestResponse
	Forbidden  *GetOfflinePaymentStatusForbiddenResponse
}

type GetOfflinePaymentStatusForbiddenResponse struct{}

type GetOfflinePaymentStatusOKResponse struct {
	Data             GetOfflinePaymentStatusOrderStatusData `json:"data,omitempty"`
	ErrorCode        string                                 `json:"error_code,omitempty"`
	ErrorDescription string                                 `json:"error_description,omitempty"`
}

type GetOfflinePaymentStatusBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOfflinePaymentStatusOrderStatusData struct {
	Status                string `json:"status,omitempty"`
	FailReasonCode        string `json:"fail_reason_code,omitempty"`
	FailReasonDescription string `json:"fail_reason_description,omitempty"`
}

type RegisterOfflineRefundRequest struct {
	Authorization string                           `json:"-"`
	ContentType   string                           `json:"-"`
	XSignature    string                           `json:"-"`
	XWBPayID      string                           `json:"-"`
	Body          RegisterOfflineRefundRequestBody `json:"-"`
}

type RegisterOfflineRefundRequestBody struct {
	Amount       int64                                  `json:"amount,omitempty"`
	CreatedAt    int64                                  `json:"created_at,omitempty"`
	CurrencyCode int64                                  `json:"currency_code,omitempty"`
	OrderID      string                                 `json:"order_id,omitempty"`
	TerminalID   string                                 `json:"terminal_id,omitempty"`
	InvoiceID    string                                 `json:"invoice_id,omitempty"`
	Positions    []RegisterOfflineRefundRequestPosition `json:"positions,omitempty"`
}

type RegisterOfflineRefundRequestPosition struct {
	Count int64  `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type RegisterOfflineRefundResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *RegisterOfflineRefundOKResponse
	BadRequest *RegisterOfflineRefundBadRequestResponse
	Forbidden  *RegisterOfflineRefundForbiddenResponse
}

type RegisterOfflineRefundForbiddenResponse struct{}

type RegisterOfflineRefundOKResponse struct {
	Data             RegisterOfflineRefundRegisterRefundData `json:"data,omitempty"`
	ErrorCode        string                                  `json:"error_code,omitempty"`
	ErrorDescription string                                  `json:"error_description,omitempty"`
}

type RegisterOfflineRefundBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type RegisterOfflineRefundRegisterRefundData struct {
	RefundID string `json:"refund_id,omitempty"`
}

type DoOfflineRefundRequest struct {
	Authorization string                     `json:"-"`
	ContentType   string                     `json:"-"`
	XWBPayID      string                     `json:"-"`
	Body          DoOfflineRefundRequestBody `json:"-"`
}

type DoOfflineRefundRequestBody struct {
	RefundID string `json:"refund_id,omitempty"`
}

type DoOfflineRefundResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *DoOfflineRefundOKResponse
	BadRequest *DoOfflineRefundBadRequestResponse
	Forbidden  *DoOfflineRefundForbiddenResponse
}

type DoOfflineRefundForbiddenResponse struct{}

type DoOfflineRefundOKResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type DoOfflineRefundBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOfflineRefundStatusRequest struct {
	RefundID      string `json:"-"`
	Authorization string `json:"-"`
	ContentType   string `json:"-"`
	XWBPayID      string `json:"-"`
}

type GetOfflineRefundStatusResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetOfflineRefundStatusOKResponse
	BadRequest *GetOfflineRefundStatusBadRequestResponse
	Forbidden  *GetOfflineRefundStatusForbiddenResponse
}

type GetOfflineRefundStatusForbiddenResponse struct{}

type GetOfflineRefundStatusOKResponse struct {
	Data             GetOfflineRefundStatusRefundStatusData `json:"data,omitempty"`
	ErrorCode        string                                 `json:"error_code,omitempty"`
	ErrorDescription string                                 `json:"error_description,omitempty"`
}

type GetOfflineRefundStatusBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetOfflineRefundStatusRefundStatusData struct {
	FailReasonCode        string `json:"fail_reason_code,omitempty"`
	FailReasonDescription string `json:"fail_reason_description,omitempty"`
	Status                string `json:"status,omitempty"`
}

type GetLoyaltyTerminalSettingsRequest struct {
	Authorization   string                                `json:"-"`
	ContentType     string                                `json:"-"`
	XRequestCountry string                                `json:"-"`
	XRequestRegion  string                                `json:"-"`
	Body            GetLoyaltyTerminalSettingsRequestBody `json:"-"`
}

type GetLoyaltyTerminalSettingsRequestBody struct {
	TerminalID string `json:"terminal_id,omitempty"`
}

type GetLoyaltyTerminalSettingsResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *GetLoyaltyTerminalSettingsOKResponse
	BadRequest *GetLoyaltyTerminalSettingsBadRequestResponse
	Forbidden  *GetLoyaltyTerminalSettingsForbiddenResponse
}

type GetLoyaltyTerminalSettingsForbiddenResponse struct{}

type GetLoyaltyTerminalSettingsOKResponse struct {
	Data             GetLoyaltyTerminalSettingsLoyaltySettingsData `json:"data,omitempty"`
	ErrorCode        string                                        `json:"error_code,omitempty"`
	ErrorDescription string                                        `json:"error_description,omitempty"`
}

type GetLoyaltyTerminalSettingsBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetLoyaltyTerminalSettingsLoyaltySettingsData struct {
	ActualDate           int64   `json:"actual_date,omitempty"`
	LoyaltyEndDate       int64   `json:"loyalty_end_date,omitempty"`
	LoyaltyRate          float64 `json:"loyalty_rate,omitempty"`
	LoyaltyStartDate     int64   `json:"loyalty_start_date,omitempty"`
	MaxCashbackAmount    int64   `json:"max_cashback_amount,omitempty"`
	MinCashbackAmount    int64   `json:"min_cashback_amount,omitempty"`
	MonthlyCashbackLimit int64   `json:"monthly_cashback_limit,omitempty"`
	Status               string  `json:"status,omitempty"`
}

type CalculateLoyaltyCashbackRequest struct {
	Authorization   string                              `json:"-"`
	ContentType     string                              `json:"-"`
	XRequestCountry string                              `json:"-"`
	XRequestRegion  string                              `json:"-"`
	Body            CalculateLoyaltyCashbackRequestBody `json:"-"`
}

type CalculateLoyaltyCashbackRequestBody struct {
	Amount       int64                                     `json:"amount,omitempty"`
	CurrencyCode int64                                     `json:"currency_code,omitempty"`
	TerminalID   string                                    `json:"terminal_id,omitempty"`
	PhoneNumber  string                                    `json:"phone_number,omitempty"`
	Positions    []CalculateLoyaltyCashbackRequestPosition `json:"positions,omitempty"`
	QRCode       string                                    `json:"qr_code,omitempty"`
}

type CalculateLoyaltyCashbackRequestPosition struct {
	Count  int64  `json:"count,omitempty"`
	Name   string `json:"name,omitempty"`
	Number int64  `json:"number,omitempty"`
	Price  int64  `json:"price,omitempty"`
}

type CalculateLoyaltyCashbackResult struct {
	StatusCode int
	Headers    http.Header
	RawBody    []byte
	OK         *CalculateLoyaltyCashbackOKResponse
	BadRequest *CalculateLoyaltyCashbackBadRequestResponse
	Forbidden  *CalculateLoyaltyCashbackForbiddenResponse
}

type CalculateLoyaltyCashbackForbiddenResponse struct{}

type CalculateLoyaltyCashbackOKResponse struct {
	Data             CalculateLoyaltyCashbackCalculateRewardData `json:"data,omitempty"`
	ErrorCode        string                                      `json:"error_code,omitempty"`
	ErrorDescription string                                      `json:"error_description,omitempty"`
}

type CalculateLoyaltyCashbackBadRequestResponse struct {
	ErrorCode        string `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type CalculateLoyaltyCashbackCalculateRewardData struct {
	TotalReward int64                                       `json:"total_reward,omitempty"`
	Positions   []CalculateLoyaltyCashbackCalculatePosition `json:"positions,omitempty"`
}

type CalculateLoyaltyCashbackCalculatePosition struct {
	RewardPerUnit   int64 `json:"reward_per_unit,omitempty"`
	TotalUnitReward int64 `json:"total_unit_reward,omitempty"`
}

const (
	WBPayValueConfirmationRejected      = "CONFIRMATION_REJECTED"
	WBPayValueConfirmationTimeExpired   = "CONFIRMATION_TIME_EXPIRED"
	WBPayValueDuplicateQrCode           = "DUPLICATE_QR_CODE"
	WBPayValueErrNone                   = "ERR_NONE"
	WBPayValueExpiredQrCode             = "EXPIRED_QR_CODE"
	WBPayValueInternalServerError       = "INTERNAL_SERVER_ERROR"
	WBPayValueInvalidPaymentToken       = "INVALID_PAYMENT_TOKEN"
	WBPayValueInvalidPhoneNumber        = "INVALID_PHONE_NUMBER"
	WBPayValueInvalidQrCode             = "INVALID_QR_CODE"
	WBPayValueInvalidRequestError       = "INVALID_REQUEST_ERROR"
	WBPayValueLimitExceeded             = "LIMIT_EXCEEDED"
	WBPayValueNotEnoughMoney            = "NOT_ENOUGH_MONEY"
	WBPayValueNotFound                  = "NOT_FOUND"
	WBPayValueNoAvailablePaymentMethods = "NO_AVAILABLE_PAYMENT_METHODS"
	WBPayValueOrderExpired              = "ORDER_EXPIRED"
	WBPayValueRefundExpired             = "REFUND_EXPIRED"
	WBPayValueRefundNotPossible         = "REFUND_NOT_POSSIBLE"
	WBPayValueRequestExpired            = "REQUEST_EXPIRED"
	WBPayValueSystemError               = "SYSTEM_ERROR"
	WBPayValueUnableToProcess           = "UNABLE_TO_PROCESS"
	WBPayValueUserNotApprove            = "USER_NOT_APPROVE"
	WBPayValueActive                    = "active"
	WBPayValueApplicationJson           = "application/json"
	WBPayValueBlocked                   = "blocked"
	WBPayValueCreated                   = "created"
	WBPayValueFailed                    = "failed"
	WBPayValueInactive                  = "inactive"
	WBPayValuePending                   = "pending"
	WBPayValueSucceeded                 = "succeeded"
)
