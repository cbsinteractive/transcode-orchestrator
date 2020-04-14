package model

type Trimming struct {
	// Defines the offset in seconds from which the encoding should start, beginning at 0.
	Offset *float64 `json:"offset,omitempty"`
	// Defines how many seconds from the input will be encoded.
	Duration *float64 `json:"duration,omitempty"`
	// If set, \"duration\" will be interpreted as a maximum and not cause an error if the input is too short
	IgnoreDurationIfInputTooShort *bool `json:"ignoreDurationIfInputTooShort,omitempty"`
	// Defines the H264 picture timing of the first frame from which the encoding should start. Any defined offset or duration in seconds will be ignored.
	StartPicTiming string `json:"startPicTiming,omitempty"`
	// Defines the H264 picture timing of the last frame, that will be included in the encoding. Any defined offset or duration in seconds will be ignored.
	EndPicTiming string `json:"endPicTiming,omitempty"`
}

