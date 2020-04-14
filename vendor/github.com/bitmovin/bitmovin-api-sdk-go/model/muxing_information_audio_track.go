package model

type MuxingInformationAudioTrack struct {
	// The stream index in the container
	Index *int32 `json:"index,omitempty"`
	// The codec used for the track
	Codec string `json:"codec,omitempty"`
	// The codec string of the track
	CodecIso string `json:"codecIso,omitempty"`
	// The bitrate of the audio track
	BitRate *int64 `json:"bitRate,omitempty"`
	Rate *int64 `json:"rate,omitempty"`
	// The sampling rate of the audio stream
	SampleRate *int32 `json:"sampleRate,omitempty"`
	// The number of channels in this audio stream
	Channels *int32 `json:"channels,omitempty"`
	// TODO add description
	Duration *float64 `json:"duration,omitempty"`
}

