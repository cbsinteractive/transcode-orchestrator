package hybrik

import (
	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

var StorageProviders = []string{"s3", "gcs", "http", "https"}

func Supported(f job.File) bool {
	p := f.Provider()
	for _, sp := range StorageProviders {
		if p == sp {
			return true
		}
	}
	return false
}

func storageBugfix(provider string, sa *hybrik.StorageAccess) *hybrik.StorageAccess {
	if provider == "gcs" {
		// Hybrik has a bug where they identify multi-region GCS -> region GCP
		// transfers as triggering egress costs, so we remove their validation for
		// GCS sources
		sa.MaxCrossRegionMB = -1
	}
	return
}

func (p *hybrikProvider) access(f job.File, creds string) *hybrik.StorageAccess {
	if creds == "" {
		if f.Provider() != "gcs" {
			return nil
		}
		creds = p.config.GCPCredentialsKey
	}
	return storageBugfix(f.Provider(), &hybrik.StorageAccess{CredentialsKey: creds})
}

func (p *hybrikProvider) location(f job.File, creds string) hybrik.TranscodeLocation {
	return hybrik.TranscodeLocation{
		StorageProvider: f.Provider(),
		Path:            f.Name,
		Access:          p.access(f, creds),
	}
}

func (p *hybrikProvider) assetURL(f job.File, creds string) hybrik.AssetURL {
	return hybrik.AssetURL{
		StorageProvider: f.Provider(),
		URL:             f.Name,
		Access:          p.access(f.Provider(), creds),
	}
}

func (p *hybrikProvider) asset(f job.File, creds string, content []hybrik.AssetContents) hybrik.AssetPayload {
	return hybrik.AssetPayload{
		StorageProvider: f.Provider(),
		URL:             f.Name,
		Contents:        content,
		Access:          p.access(f, creds),
	}
}
