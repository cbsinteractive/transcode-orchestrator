package model
type AacChannelLayout string

// List of AacChannelLayout
const (
	AacChannelLayout_NONE AacChannelLayout = "NONE"
	AacChannelLayout_MONO AacChannelLayout = "MONO"
	AacChannelLayout_CL_STEREO AacChannelLayout = "STEREO"
	AacChannelLayout_CL_SURROUND AacChannelLayout = "SURROUND"
	AacChannelLayout_CL_4_0 AacChannelLayout = "4.0"
	AacChannelLayout_CL_5_0_BACK AacChannelLayout = "5.0_BACK"
	AacChannelLayout_CL_5_1_BACK AacChannelLayout = "5.1_BACK"
	AacChannelLayout_CL_7_1 AacChannelLayout = "7.1"
	AacChannelLayout_CL_7_1_WIDE_BACK AacChannelLayout = "7.1_WIDE_BACK"
)