package model

type AutoRestartConfiguration struct {
	// If no segments were generated for the given number of seconds, a restart is triggered. Minimum: 30.0
	SegmentsWrittenTimeout *float64 `json:"segmentsWrittenTimeout,omitempty"`
	// If no data was written for the given number of seconds, a restart is triggered. Minimum: 30.0
	BytesWrittenTimeout *float64 `json:"bytesWrittenTimeout,omitempty"`
	// If no frames were generated for the given number of seconds, a restart is triggered. Minimum: 30.0
	FramesWrittenTimeout *float64 `json:"framesWrittenTimeout,omitempty"`
	// If HLS manifests were not updated for the given number of seconds, a restart is triggered. Minimum: 30.0
	HlsManifestsUpdateTimeout *float64 `json:"hlsManifestsUpdateTimeout,omitempty"`
	// If DASH manifests were not updated for the given number of seconds, a restart is triggered. Minimum: 30.0
	DashManifestsUpdateTimeout *float64 `json:"dashManifestsUpdateTimeout,omitempty"`
	// Defines a schedule for restarts using the unix crontab syntax. This example would trigger a restart every monday at 05:30 (AM)
	ScheduleExpression string `json:"scheduleExpression,omitempty"`
}

