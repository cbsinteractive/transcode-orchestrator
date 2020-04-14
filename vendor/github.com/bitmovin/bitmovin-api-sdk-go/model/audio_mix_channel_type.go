package model
// AudioMixChannelType : Type of this channel
type AudioMixChannelType string

// List of AudioMixChannelType
const (
	AudioMixChannelType_CHANNEL_NUMBER AudioMixChannelType = "CHANNEL_NUMBER"
	AudioMixChannelType_FRONT_LEFT AudioMixChannelType = "FRONT_LEFT"
	AudioMixChannelType_FRONT_RIGHT AudioMixChannelType = "FRONT_RIGHT"
	AudioMixChannelType_CENTER AudioMixChannelType = "CENTER"
	AudioMixChannelType_LOW_FREQUENCY AudioMixChannelType = "LOW_FREQUENCY"
	AudioMixChannelType_BACK_LEFT AudioMixChannelType = "BACK_LEFT"
	AudioMixChannelType_BACK_RIGHT AudioMixChannelType = "BACK_RIGHT"
	AudioMixChannelType_SURROUND_LEFT AudioMixChannelType = "SURROUND_LEFT"
	AudioMixChannelType_SURROUND_RIGHT AudioMixChannelType = "SURROUND_RIGHT"
)