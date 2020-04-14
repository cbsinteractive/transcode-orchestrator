package model
type VorbisChannelLayout string

// List of VorbisChannelLayout
const (
	VorbisChannelLayout_NONE VorbisChannelLayout = "NONE"
	VorbisChannelLayout_MONO VorbisChannelLayout = "MONO"
	VorbisChannelLayout_CL_STEREO VorbisChannelLayout = "STEREO"
	VorbisChannelLayout_CL_SURROUND VorbisChannelLayout = "SURROUND"
	VorbisChannelLayout_CL_QUAD VorbisChannelLayout = "QUAD"
	VorbisChannelLayout_CL_2_1 VorbisChannelLayout = "2.1"
	VorbisChannelLayout_CL_2_2 VorbisChannelLayout = "2.2"
	VorbisChannelLayout_CL_3_1 VorbisChannelLayout = "3.1"
	VorbisChannelLayout_CL_4_0 VorbisChannelLayout = "4.0"
	VorbisChannelLayout_CL_5_1 VorbisChannelLayout = "5.1"
	VorbisChannelLayout_CL_5_1_BACK VorbisChannelLayout = "5.1_BACK"
)