package model
// DeinterlaceMode : Specifies the method how fields are converted to frames
type DeinterlaceMode string

// List of DeinterlaceMode
const (
	DeinterlaceMode_FRAME DeinterlaceMode = "FRAME"
	DeinterlaceMode_FIELD DeinterlaceMode = "FIELD"
	DeinterlaceMode_FRAME_NOSPATIAL DeinterlaceMode = "FRAME_NOSPATIAL"
	DeinterlaceMode_FIELD_NOSPATIAL DeinterlaceMode = "FIELD_NOSPATIAL"
)