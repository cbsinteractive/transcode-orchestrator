package model
import (
	"time"
)

type ModelTask struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Current status (required)
	Status Status `json:"status,omitempty"`
	// Estimated ETA in seconds
	Eta *float64 `json:"eta,omitempty"`
	// Progress in percent
	Progress *int32 `json:"progress,omitempty"`
	// List of subtasks
	Subtasks []Subtask `json:"subtasks,omitempty"`
	// Task specific messages
	Messages []Message `json:"messages,omitempty"`
	// Timestamp when the task was created, formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp when the task status changed to \"QUEUED\", formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	QueuedAt *time.Time `json:"queuedAt,omitempty"`
	// Timestamp when the task status changed to to \"RUNNING\", formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	RunningAt *time.Time `json:"runningAt,omitempty"`
	// Timestamp when the task status changed to \"FINISHED\", formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Timestamp when the task status changed to \"ERROR\", formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ErrorAt *time.Time `json:"errorAt,omitempty"`
	// Additional optional error details
	Error *ErrorDetails `json:"error,omitempty"`
}

