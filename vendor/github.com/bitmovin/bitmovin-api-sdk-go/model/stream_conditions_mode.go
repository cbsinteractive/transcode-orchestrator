package model
type StreamConditionsMode string

// List of StreamConditionsMode
const (
	StreamConditionsMode_DROP_MUXING StreamConditionsMode = "DROP_MUXING"
	StreamConditionsMode_DROP_STREAM StreamConditionsMode = "DROP_STREAM"
)