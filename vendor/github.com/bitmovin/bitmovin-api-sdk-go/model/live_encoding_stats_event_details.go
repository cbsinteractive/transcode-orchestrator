package model

type LiveEncodingStatsEventDetails struct {
	EventName LiveEncodingEventName `json:"eventName,omitempty"`
	// The Audio/Video Drift in seconds. The drift was corrected by the RESYNCING event (occurs at event: RESYNCING)
	AvDriftInSeconds *int32 `json:"avDriftInSeconds,omitempty"`
	// The time the stream was in idle state in seconds (occurs at event: IDLE)
	IdleDurationInSeconds *int32 `json:"idleDurationInSeconds,omitempty"`
	// An optional error message, when the event is in error state (occurs at event: ERROR)
	ErrorMessage string `json:"errorMessage,omitempty"`
}

