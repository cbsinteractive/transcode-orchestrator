package container

import "github.com/cbsinteractive/transcode-orchestrator/provider"

// StatusEnricher enriches status information for output containers
type StatusEnricher interface {
	Enrich(provider.JobStatus) (provider.JobStatus, error)
}
