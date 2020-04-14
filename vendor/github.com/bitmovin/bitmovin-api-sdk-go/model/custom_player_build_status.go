package model

type CustomPlayerBuildStatus struct {
	// Status of the player build (required)
	Status Status `json:"status,omitempty"`
	// The estimated time span of the custom player build in seconds.
	Eta *int32 `json:"eta,omitempty"`
	// The actual progress of the custom player build. (required)
	Progress *int32 `json:"progress,omitempty"`
	Messages *Message `json:"messages,omitempty"`
	Subtasks string `json:"subtasks,omitempty"`
}

