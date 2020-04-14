package model
import (
	"time"
)

type EmailNotification struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Notify when condition resolves after it was met
	Resolve *bool `json:"resolve,omitempty"`
	// Specific resource, e.g. encoding id
	ResourceId string `json:"resourceId,omitempty"`
	// Last time the notification was triggered
	TriggeredAt *time.Time `json:"triggeredAt,omitempty"`
	Type string `json:"type,omitempty"`
	EventType string `json:"eventType,omitempty"`
	Category string `json:"category,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Muted *bool `json:"muted,omitempty"`
	Emails []string `json:"emails,omitempty"`
}

