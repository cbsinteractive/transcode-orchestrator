package model

type VideoStream struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Position starts from 0 and indicates the position of the stream in the media. 0 means that this is the first stream found in the media
	Position *int32 `json:"position,omitempty"`
	// Duration of the stream in seconds
	Duration *float64 `json:"duration,omitempty"`
	// Codec of the stream
	Codec string `json:"codec,omitempty"`
	// Frame rate of the video
	Fps string `json:"fps,omitempty"`
	// Bitrate in bps
	Bitrate string `json:"bitrate,omitempty"`
	// Bitrate in bps (the same as bitrate, but represented as an integer value)
	Rate *int64 `json:"rate,omitempty"`
	// Width of the video (required)
	Width *int32 `json:"width,omitempty"`
	// Height of the video (required)
	Height *int32 `json:"height,omitempty"`
	// Pixel aspect ratio of the video. Default is 1.0
	Par *float64 `json:"par,omitempty"`
	// Rotation of the video for mobile devices. Default is 0.
	Rotation *int32 `json:"rotation,omitempty"`
}

