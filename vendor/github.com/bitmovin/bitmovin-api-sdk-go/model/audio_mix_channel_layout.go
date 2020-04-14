package model
// AudioMixChannelLayout : Channel layout of this audio mix
type AudioMixChannelLayout string

// List of AudioMixChannelLayout
const (
	AudioMixChannelLayout_NONE AudioMixChannelLayout = "NONE"
	AudioMixChannelLayout_CL_MONO AudioMixChannelLayout = "MONO"
	AudioMixChannelLayout_CL_STEREO AudioMixChannelLayout = "STEREO"
	AudioMixChannelLayout_CL_SURROUND AudioMixChannelLayout = "SURROUND"
	AudioMixChannelLayout_CL_QUAD AudioMixChannelLayout = "QUAD"
	AudioMixChannelLayout_CL_OCTAGONAL AudioMixChannelLayout = "OCTAGONAL"
	AudioMixChannelLayout_CL_HEXAGONAL AudioMixChannelLayout = "HEXAGONAL"
	AudioMixChannelLayout_CL_STEREO_DOWNMIX AudioMixChannelLayout = "STEREO_DOWNMIX"
	AudioMixChannelLayout_CL_2_1 AudioMixChannelLayout = "2.1"
	AudioMixChannelLayout_CL_2_2 AudioMixChannelLayout = "2.2"
	AudioMixChannelLayout_CL_3_1 AudioMixChannelLayout = "3.1"
	AudioMixChannelLayout_CL_4_0 AudioMixChannelLayout = "4.0"
	AudioMixChannelLayout_CL_4_1 AudioMixChannelLayout = "4.1"
	AudioMixChannelLayout_CL_5_0 AudioMixChannelLayout = "5.0"
	AudioMixChannelLayout_CL_5_1 AudioMixChannelLayout = "5.1"
	AudioMixChannelLayout_CL_5_0_BACK AudioMixChannelLayout = "5.0_BACK"
	AudioMixChannelLayout_CL_5_1_BACK AudioMixChannelLayout = "5.1_BACK"
	AudioMixChannelLayout_CL_6_0 AudioMixChannelLayout = "6.0"
	AudioMixChannelLayout_CL_6_0_FRONT AudioMixChannelLayout = "6.0_FRONT"
	AudioMixChannelLayout_CL_6_1 AudioMixChannelLayout = "6.1"
	AudioMixChannelLayout_CL_6_1_BACK AudioMixChannelLayout = "6.1_BACK"
	AudioMixChannelLayout_CL_6_1_FRONT AudioMixChannelLayout = "6.1_FRONT"
	AudioMixChannelLayout_CL_7_0 AudioMixChannelLayout = "7.0"
	AudioMixChannelLayout_CL_7_0_FRONT AudioMixChannelLayout = "7.0_FRONT"
	AudioMixChannelLayout_CL_7_1 AudioMixChannelLayout = "7.1"
	AudioMixChannelLayout_CL_7_1_WIDE AudioMixChannelLayout = "7.1_WIDE"
	AudioMixChannelLayout_CL_7_1_WIDE_BACK AudioMixChannelLayout = "7.1_WIDE_BACK"
)