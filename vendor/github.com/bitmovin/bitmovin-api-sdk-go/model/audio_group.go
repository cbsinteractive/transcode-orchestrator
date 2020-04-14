package model

type AudioGroup struct {
	// Name of the audio group (required)
	Name string `json:"name,omitempty"`
	// Priority of the audio group (required)
	Priority *int32 `json:"priority,omitempty"`
}

