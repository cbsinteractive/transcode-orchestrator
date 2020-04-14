package model

type AudioMixChannel struct {
	// Channel number of this mix (starting with 0) (required)
	ChannelNumber *int32 `json:"channelNumber,omitempty"`
	// List of source channels to be mixed (required)
	SourceChannels []SourceChannel `json:"sourceChannels,omitempty"`
}

