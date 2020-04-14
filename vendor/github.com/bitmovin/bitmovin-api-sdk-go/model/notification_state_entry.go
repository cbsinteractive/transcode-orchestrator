package model
import (
	"time"
)

type NotificationStateEntry struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	State NotificationStates `json:"state,omitempty"`
	// Indicate if notification was sent (required)
	Muted *bool `json:"muted,omitempty"`
	// The notification this state belongs to (required)
	NotificationId string `json:"notificationId,omitempty"`
	// Indicate if triggered for specific resource (required)
	ResourceId string `json:"resourceId,omitempty"`
	TriggeredAt *time.Time `json:"triggeredAt,omitempty"`
}

