package model

type ObjectDetectionTimestampResult struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Time in seconds where the object was detected in the video
	Timestamp *float64 `json:"timestamp,omitempty"`
	// Objects detected for the given timestamp
	Objects []ObjectDetectionResult `json:"objects,omitempty"`
}

