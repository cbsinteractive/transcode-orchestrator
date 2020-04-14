package model

type PerTitle struct {
	// Per-Title configuration for H264
	H264Configuration *H264PerTitleConfiguration `json:"h264Configuration,omitempty"`
	// Per-Title configuration for H265
	H265Configuration *H265PerTitleConfiguration `json:"h265Configuration,omitempty"`
	// Per-Title configuration for VP9
	Vp9Configuration *Vp9PerTitleConfiguration `json:"vp9Configuration,omitempty"`
}

