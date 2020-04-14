package model
// InputColorRange : Override the color range detected in the input file. If not set the input color range will be automatically detected if possible.
type InputColorRange string

// List of InputColorRange
const (
	InputColorRange_UNSPECIFIED InputColorRange = "UNSPECIFIED"
	InputColorRange_MPEG InputColorRange = "MPEG"
	InputColorRange_JPEG InputColorRange = "JPEG"
)