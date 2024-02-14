package openrouter

type ErrStatusCode int

var (
	// List of error status codes
	// Taken from https://openrouter.ai/docs#errors

	ErrBadRequest         ErrStatusCode = 400 // Invalid or missing parameters, CORS
	ErrUnauthorized       ErrStatusCode = 401 // Invalid Credentials
	ErrPaymentRequired    ErrStatusCode = 402 // Your account or API Key has insufficient credits
	ErrForbidden          ErrStatusCode = 403 // Your chosen model requires moderation and your input was flagged
	ErrTimeout            ErrStatusCode = 408 // The request timed out
	ErrRateLimit          ErrStatusCode = 429 // You have exceeded your rate limit
	ErrBadGateway         ErrStatusCode = 502 // Your chosen model is down or we received an invalid response from it
	ErrServiceUnavailable ErrStatusCode = 503 // There is no available model provider that meets your routing requirements
)

type ErrorResponse struct {
	Code     int                    `json:"code"`
	Msg      string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata"`
}
