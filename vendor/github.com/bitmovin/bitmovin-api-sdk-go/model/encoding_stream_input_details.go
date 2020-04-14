package model

type EncodingStreamInputDetails struct {
	// Format name
	FormatName string `json:"formatName,omitempty"`
	// The start time in seconds
	StartTime *float64 `json:"startTime,omitempty"`
	// Duration in seconds
	Duration *float64 `json:"duration,omitempty"`
	// Input file size in bytes
	Size *int64 `json:"size,omitempty"`
	// Bitrate in bps
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Additional metadata saved in the input file
	Tags *map[string]map[string]interface{} `json:"tags,omitempty"`
	// Video streams in the input file
	VideoStreams []VideoStream `json:"videoStreams,omitempty"`
	// Audio stream in the input file
	AudioStreams []AudioStream `json:"audioStreams,omitempty"`
	// Meta data streams in the input file
	MetaStreams []MediaStream `json:"metaStreams,omitempty"`
	// Subtitle streams in the input file
	SubtitleStreams []SubtitleStream `json:"subtitleStreams,omitempty"`
}

