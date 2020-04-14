package model

type ErrorDetails struct {
	// Specific error code (required)
	Code *int32 `json:"code,omitempty"`
	// Error group name (required)
	Category string `json:"category,omitempty"`
	// Detailed error message (required)
	Text string `json:"text,omitempty"`
	// Information if the encoding could potentially succeed when retrying. (required)
	RetryHint RetryHint `json:"retryHint,omitempty"`
}

