package model
type ErrorRetryHint string

// List of ErrorRetryHint
const (
	ErrorRetryHint_RETRY ErrorRetryHint = "RETRY"
	ErrorRetryHint_NO_RETRY ErrorRetryHint = "NO_RETRY"
	ErrorRetryHint_RETRY_IN_DIFFERENT_REGION ErrorRetryHint = "RETRY_IN_DIFFERENT_REGION"
)