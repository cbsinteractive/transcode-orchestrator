package model
type MessageType string

// List of MessageType
const (
	MessageType_ERROR MessageType = "ERROR"
	MessageType_WARNING MessageType = "WARNING"
	MessageType_INFO MessageType = "INFO"
	MessageType_DEBUG MessageType = "DEBUG"
	MessageType_TRACE MessageType = "TRACE"
)