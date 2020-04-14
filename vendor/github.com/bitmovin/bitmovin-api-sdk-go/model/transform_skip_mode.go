package model
// TransformSkipMode : Enable evaluation of transform skip (bypass DCT but still use quantization) coding for 4x4 TU coded blocks.
type TransformSkipMode string

// List of TransformSkipMode
const (
	TransformSkipMode_NONE TransformSkipMode = "NONE"
	TransformSkipMode_NORMAL TransformSkipMode = "NORMAL"
	TransformSkipMode_FAST TransformSkipMode = "FAST"
)