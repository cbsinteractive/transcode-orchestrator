package model

// The segmenting configuration for this WebVTT sidecar file. If this is set, the given vtt file will be chunked.
type WebVttSidecarFileSegmentation struct {
	// The length of the WebVTT fragments in seconds (required)
	SegmentLength *float64 `json:"segmentLength,omitempty"`
}

