package model
type LiveEncodingEventName string

// List of LiveEncodingEventName
const (
	LiveEncodingEventName_FIRST_CONNECT LiveEncodingEventName = "FIRST_CONNECT"
	LiveEncodingEventName_DISCONNECT LiveEncodingEventName = "DISCONNECT"
	LiveEncodingEventName_RECONNECT LiveEncodingEventName = "RECONNECT"
	LiveEncodingEventName_RESYNCING LiveEncodingEventName = "RESYNCING"
	LiveEncodingEventName_IDLE LiveEncodingEventName = "IDLE"
	LiveEncodingEventName_ERROR LiveEncodingEventName = "ERROR"
)