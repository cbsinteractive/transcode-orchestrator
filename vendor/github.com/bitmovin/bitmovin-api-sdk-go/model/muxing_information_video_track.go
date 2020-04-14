package model

type MuxingInformationVideoTrack struct {
	// The stream index in the container
	Index *int32 `json:"index,omitempty"`
	// The codec used for the track
	Codec string `json:"codec,omitempty"`
	// The codec string of the track
	CodecIso string `json:"codecIso,omitempty"`
	// The bitrate of the video track
	BitRate *int64 `json:"bitRate,omitempty"`
	// TODO add description
	Rate *int64 `json:"rate,omitempty"`
	// The pixel format used
	PixelFormat string `json:"pixelFormat,omitempty"`
	// The frame mode (e.g. progressive)
	FrameMode string `json:"frameMode,omitempty"`
	// The width of the frame in pixel
	FrameWidth *int32 `json:"frameWidth,omitempty"`
	// The height of the frame in pixel
	FrameHeight *int32 `json:"frameHeight,omitempty"`
	// The frame rate of the stream in fractional format
	FrameRate string `json:"frameRate,omitempty"`
	// The start time in seconds
	StartTime *float64 `json:"startTime,omitempty"`
	// The duration in seconds
	Duration *float64 `json:"duration,omitempty"`
	// The number of frames of that video track
	NumberOfFrames *int64 `json:"numberOfFrames,omitempty"`
	// The aspect ratio of the stream
	AspectRatio string `json:"aspectRatio,omitempty"`
}

