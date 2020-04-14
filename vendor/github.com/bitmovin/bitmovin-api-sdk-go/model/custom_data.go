package model
import (
	"time"
)

type CustomData struct {
	// User-specific meta data. This can hold a custom JSON object.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

