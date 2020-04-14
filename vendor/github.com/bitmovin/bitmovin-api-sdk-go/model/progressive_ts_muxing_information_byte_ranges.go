package model

type ProgressiveTsMuxingInformationByteRanges struct {
	// Number of the segment (starting at 0) (required)
	SegmentNumber *int32 `json:"segmentNumber,omitempty"`
	// The position of the first byte of the segment
	StartBytes *int64 `json:"startBytes,omitempty"`
	// The position of the last byte of the segment
	EndBytes *int64 `json:"endBytes,omitempty"`
	// The duration of the segment in seconds
	Duration *float64 `json:"duration,omitempty"`
}

