package model

type Period struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Starting time in seconds
	Start *float64 `json:"start,omitempty"`
	// Duration in seconds
	Duration *float64 `json:"duration,omitempty"`
}

