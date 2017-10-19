package model

// ErrorResponse describes a typical error response from our service.
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
