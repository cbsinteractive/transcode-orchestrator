package model

type ResourceLimitContainer struct {
	Resource ResourceType `json:"resource,omitempty"`
	Limits []ResourceLimit `json:"limits,omitempty"`
}

