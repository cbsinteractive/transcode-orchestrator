package model
// InputColorSpace : Override the color space detected in the input file. If not set the input color space will be automatically detected if possible.
type InputColorSpace string

// List of InputColorSpace
const (
	InputColorSpace_UNSPECIFIED InputColorSpace = "UNSPECIFIED"
	InputColorSpace_RGB InputColorSpace = "RGB"
	InputColorSpace_BT709 InputColorSpace = "BT709"
	InputColorSpace_FCC InputColorSpace = "FCC"
	InputColorSpace_BT470BG InputColorSpace = "BT470BG"
	InputColorSpace_SMPTE170M InputColorSpace = "SMPTE170M"
	InputColorSpace_SMPTE240M InputColorSpace = "SMPTE240M"
	InputColorSpace_YCGCO InputColorSpace = "YCGCO"
	InputColorSpace_YCOCG InputColorSpace = "YCOCG"
	InputColorSpace_BT2020_NCL InputColorSpace = "BT2020_NCL"
	InputColorSpace_BT2020_CL InputColorSpace = "BT2020_CL"
	InputColorSpace_SMPTE2085 InputColorSpace = "SMPTE2085"
)