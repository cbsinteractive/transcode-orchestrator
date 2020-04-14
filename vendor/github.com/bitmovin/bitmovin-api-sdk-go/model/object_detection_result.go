package model

type ObjectDetectionResult struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Name of the object that has been detected (in English)
	DetectedObject string `json:"detectedObject,omitempty"`
	// Time in seconds where the object was detected in the video
	Timestamp *float64 `json:"timestamp,omitempty"`
	// A number between 0 and 1 indicating the confidence of the detection
	Score *float64 `json:"score,omitempty"`
	// A box indicating the position and size of the detected object within the frame
	BoundingBox *ObjectDetectionBoundingBox `json:"boundingBox,omitempty"`
}

