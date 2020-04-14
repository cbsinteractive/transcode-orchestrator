package model
// InterlaceMode : How to interleave the input frames
type InterlaceMode string

// List of InterlaceMode
const (
	InterlaceMode_TOP InterlaceMode = "TOP"
	InterlaceMode_BOTTOM InterlaceMode = "BOTTOM"
	InterlaceMode_DROP_EVEN InterlaceMode = "DROP_EVEN"
	InterlaceMode_DROP_ODD InterlaceMode = "DROP_ODD"
	InterlaceMode_PAD InterlaceMode = "PAD"
	InterlaceMode_INTERLACE_X2 InterlaceMode = "INTERLACE_X2"
	InterlaceMode_MERGE InterlaceMode = "MERGE"
	InterlaceMode_MERGE_X2 InterlaceMode = "MERGE_X2"
)