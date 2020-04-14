package model

type RtmpIngestPoint struct {
	// The name of the application where the ingest is streamed to. This has to be unique for each ingest point (required)
	ApplicationName string `json:"applicationName,omitempty"`
	// The stream key for the backup input (required)
	StreamKey string `json:"streamKey,omitempty"`
}

