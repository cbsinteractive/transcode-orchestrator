package model
// H264Trellis : Enables or disables Trellis quantization. NOTE: This requires cabac
type H264Trellis string

// List of H264Trellis
const (
	H264Trellis_DISABLED H264Trellis = "DISABLED"
	H264Trellis_ENABLED_FINAL_MB H264Trellis = "ENABLED_FINAL_MB"
	H264Trellis_ENABLED_ALL H264Trellis = "ENABLED_ALL"
)