package model

type KubernetesClusterConfiguration struct {
	// Number of parallel scheduled encodings on the Kubernetes cluster (required)
	ParallelEncodings *int32 `json:"parallelEncodings,omitempty"`
	// Number of worker nodes used for each encoding on the Kubernetes cluster (required)
	WorkersPerEncoding *int32 `json:"workersPerEncoding,omitempty"`
}

