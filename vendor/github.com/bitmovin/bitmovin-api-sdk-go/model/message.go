package model
import (
	"time"
)

type Message struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Message type giving a hint on the importance of the message (log level) (required)
	Type MessageType `json:"type,omitempty"`
	// Message text (required)
	Text string `json:"text,omitempty"`
	// Name of the field to which the message is referring to
	Field string `json:"field,omitempty"`
	// collection of links to webpages containing further information on the topic
	Links []Link `json:"links,omitempty"`
	// Service-specific information
	More *map[string]interface{} `json:"more,omitempty"`
	// Timestamp when the message occured
	Date *time.Time `json:"date,omitempty"`
}

