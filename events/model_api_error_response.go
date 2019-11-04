package events

// APIErrorResponse struct
type APIErrorResponse struct {
	Errors []APIErrorResponseErrors `json:"errors,omitempty"`
}
