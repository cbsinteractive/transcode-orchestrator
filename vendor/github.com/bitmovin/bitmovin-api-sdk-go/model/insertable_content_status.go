package model
type InsertableContentStatus string

// List of InsertableContentStatus
const (
	InsertableContentStatus_CREATED InsertableContentStatus = "CREATED"
	InsertableContentStatus_DOWNLOADING InsertableContentStatus = "DOWNLOADING"
	InsertableContentStatus_READY InsertableContentStatus = "READY"
	InsertableContentStatus_ERROR InsertableContentStatus = "ERROR"
)