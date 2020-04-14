package model
type Status string

// List of Status
const (
	Status_CREATED Status = "CREATED"
	Status_QUEUED Status = "QUEUED"
	Status_RUNNING Status = "RUNNING"
	Status_FINISHED Status = "FINISHED"
	Status_ERROR Status = "ERROR"
	Status_CANCELED Status = "CANCELED"
)