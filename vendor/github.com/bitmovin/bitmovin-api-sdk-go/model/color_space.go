package model
// ColorSpace : The color space to be applied
type ColorSpace string

// List of ColorSpace
const (
	ColorSpace_UNSPECIFIED ColorSpace = "UNSPECIFIED"
	ColorSpace_RGB ColorSpace = "RGB"
	ColorSpace_BT709 ColorSpace = "BT709"
	ColorSpace_FCC ColorSpace = "FCC"
	ColorSpace_BT470BG ColorSpace = "BT470BG"
	ColorSpace_SMPTE170M ColorSpace = "SMPTE170M"
	ColorSpace_SMPTE240M ColorSpace = "SMPTE240M"
	ColorSpace_YCGCO ColorSpace = "YCGCO"
	ColorSpace_YCOCG ColorSpace = "YCOCG"
	ColorSpace_BT2020_NCL ColorSpace = "BT2020_NCL"
	ColorSpace_BT2020_CL ColorSpace = "BT2020_CL"
	ColorSpace_SMPTE2085 ColorSpace = "SMPTE2085"
)