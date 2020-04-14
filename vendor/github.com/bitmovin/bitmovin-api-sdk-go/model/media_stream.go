package model

type MediaStream struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Position starts from 0 and indicates the position of the stream in the media. 0 means that this is the first stream found in the media
	Position *int32 `json:"position,omitempty"`
	// Duration of the stream in seconds
	Duration *float64 `json:"duration,omitempty"`
	// Codec of the stream
	Codec string `json:"codec,omitempty"`
}

