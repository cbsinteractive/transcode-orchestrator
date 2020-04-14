package model

type SourceChannel struct {
	// Gain for this source channel. Default is 1.0.
	Gain *float64 `json:"gain,omitempty"`
	Type SourceChannelType `json:"type,omitempty"`
	// Number of this source channel. If type is 'CHANNEL_NUMBER', this must be set.
	ChannelNumber *int32 `json:"channelNumber,omitempty"`
}

