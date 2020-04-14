package model

type AudioStream struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Position of the stream in the file
	Position *int32 `json:"position,omitempty"`
	// Duration of the stream in seconds
	Duration *float64 `json:"duration,omitempty"`
	// Codec of the stream
	Codec string `json:"codec,omitempty"`
	// Audio sampling rate in Hz
	SampleRate *int32 `json:"sampleRate,omitempty"`
	// Bitrate in bps
	Bitrate string `json:"bitrate,omitempty"`
	// Bitrate in bps
	Rate *int64 `json:"rate,omitempty"`
	// Audio channel format
	ChannelFormat string `json:"channelFormat,omitempty"`
	// Language of the stream
	Language string `json:"language,omitempty"`
	// Hearing impaired support
	HearingImpaired *bool `json:"hearingImpaired,omitempty"`
}

