package model
type EncodingMode string

// List of EncodingMode
const (
	EncodingMode_STANDARD EncodingMode = "STANDARD"
	EncodingMode_SINGLE_PASS EncodingMode = "SINGLE_PASS"
	EncodingMode_TWO_PASS EncodingMode = "TWO_PASS"
	EncodingMode_THREE_PASS EncodingMode = "THREE_PASS"
)