package model
// H264InterlaceMode : Using TOP_FIELD_FIRST or BOTTOM_FIELD_FIRST will output interlaced video
type H264InterlaceMode string

// List of H264InterlaceMode
const (
	H264InterlaceMode_NONE H264InterlaceMode = "NONE"
	H264InterlaceMode_TOP_FIELD_FIRST H264InterlaceMode = "TOP_FIELD_FIRST"
	H264InterlaceMode_BOTTOM_FIELD_FIRST H264InterlaceMode = "BOTTOM_FIELD_FIRST"
)