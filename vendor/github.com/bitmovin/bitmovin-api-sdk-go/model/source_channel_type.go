package model
// SourceChannelType : Type of this source channel
type SourceChannelType string

// List of SourceChannelType
const (
	SourceChannelType_CHANNEL_NUMBER SourceChannelType = "CHANNEL_NUMBER"
	SourceChannelType_FRONT_LEFT SourceChannelType = "FRONT_LEFT"
	SourceChannelType_FRONT_RIGHT SourceChannelType = "FRONT_RIGHT"
	SourceChannelType_CENTER SourceChannelType = "CENTER"
	SourceChannelType_LOW_FREQUENCY SourceChannelType = "LOW_FREQUENCY"
	SourceChannelType_BACK_LEFT SourceChannelType = "BACK_LEFT"
	SourceChannelType_BACK_RIGHT SourceChannelType = "BACK_RIGHT"
	SourceChannelType_SURROUND_LEFT SourceChannelType = "SURROUND_LEFT"
	SourceChannelType_SURROUND_RIGHT SourceChannelType = "SURROUND_RIGHT"
)