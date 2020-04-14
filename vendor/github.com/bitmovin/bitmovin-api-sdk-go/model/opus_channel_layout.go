package model
type OpusChannelLayout string

// List of OpusChannelLayout
const (
	OpusChannelLayout_NONE OpusChannelLayout = "NONE"
	OpusChannelLayout_MONO OpusChannelLayout = "MONO"
	OpusChannelLayout_CL_STEREO OpusChannelLayout = "STEREO"
	OpusChannelLayout_CL_SURROUND OpusChannelLayout = "SURROUND"
	OpusChannelLayout_CL_QUAD OpusChannelLayout = "QUAD"
	OpusChannelLayout_CL_4_1 OpusChannelLayout = "4.1"
	OpusChannelLayout_CL_5_0_BACK OpusChannelLayout = "5.0_BACK"
	OpusChannelLayout_XCL_5_1_BACK OpusChannelLayout = "5.1_BACK"
	OpusChannelLayout_CL_6_1 OpusChannelLayout = "6.1"
	OpusChannelLayout_CL_7_1 OpusChannelLayout = "7.1"
)