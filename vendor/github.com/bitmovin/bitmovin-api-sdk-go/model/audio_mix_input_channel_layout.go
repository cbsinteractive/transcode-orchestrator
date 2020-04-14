package model
type AudioMixInputChannelLayout string

// List of AudioMixInputChannelLayout
const (
	AudioMixInputChannelLayout_NONE AudioMixInputChannelLayout = "NONE"
	AudioMixInputChannelLayout_MONO AudioMixInputChannelLayout = "MONO"
	AudioMixInputChannelLayout_CL_STEREO AudioMixInputChannelLayout = "STEREO"
	AudioMixInputChannelLayout_CL_SURROUND AudioMixInputChannelLayout = "SURROUND"
	AudioMixInputChannelLayout_CL_4_0 AudioMixInputChannelLayout = "4.0"
	AudioMixInputChannelLayout_CL_5_0_BACK AudioMixInputChannelLayout = "5.0_BACK"
	AudioMixInputChannelLayout_CL_5_1_BACK AudioMixInputChannelLayout = "5.1_BACK"
	AudioMixInputChannelLayout_CL_7_1 AudioMixInputChannelLayout = "7.1"
	AudioMixInputChannelLayout_CL_7_1_WIDE_BACK AudioMixInputChannelLayout = "7.1_WIDE_BACK"
)