package model

type InfrastructureSettings struct {
	// Id of a custom infrastructure, e.g., Kubernetes Cluster
	InfrastructureId string `json:"infrastructureId,omitempty"`
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
}

