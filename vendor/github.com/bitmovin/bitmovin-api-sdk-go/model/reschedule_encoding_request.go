package model

type RescheduleEncodingRequest struct {
	// Id of a custom infrastructure, e.g., Kubernetes Cluster
	InfrastructureId string `json:"infrastructureId,omitempty"`
}

