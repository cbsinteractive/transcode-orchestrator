package model

type LiveEncodingStats struct {
	Status LiveEncodingStatus `json:"status,omitempty"`
	// List of events
	Events []LiveEncodingStatsEvent `json:"events,omitempty"`
	// List of statistics
	Statistics []StreamInfos `json:"statistics,omitempty"`
}

