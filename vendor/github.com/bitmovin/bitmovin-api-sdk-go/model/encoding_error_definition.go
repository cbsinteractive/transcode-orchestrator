package model

type EncodingErrorDefinition struct {
	// The error code.
	Code *int32 `json:"code,omitempty"`
	// The error category.
	Category string `json:"category,omitempty"`
	// The error message, optional. Can include placeholders like {1}, {2} which are replaced with the respective dependency when the error is thrown.
	Message string `json:"message,omitempty"`
	// The returned error description.
	Description string `json:"description,omitempty"`
	// Indicates if the call that caused the error should be retried.
	RetryHint ErrorRetryHint `json:"retryHint,omitempty"`
}

