package model
type DecodingErrorMode string

// List of DecodingErrorMode
const (
	DecodingErrorMode_FAIL_ON_ERROR DecodingErrorMode = "FAIL_ON_ERROR"
	DecodingErrorMode_DUPLICATE_FRAMES DecodingErrorMode = "DUPLICATE_FRAMES"
)