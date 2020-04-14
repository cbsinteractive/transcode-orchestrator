package model
type RetryHint string

// List of RetryHint
const (
	RetryHint_RETRY RetryHint = "RETRY"
	RetryHint_RETRY_IN_DIFFERENT_REGION RetryHint = "RETRY_IN_DIFFERENT_REGION"
	RetryHint_NO_RETRY RetryHint = "NO_RETRY"
)