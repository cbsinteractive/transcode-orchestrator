package model

type StreamFilterList struct {
	// List of stream filters (required)
	Filters []StreamFilter `json:"filters,omitempty"`
}

