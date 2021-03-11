package hybrik

import (
	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
)

type storageProvider string

func (p storageProvider) supportsSegmentedRendering() bool { return p != storageProviderHTTP }
func (p storageProvider) string() string                   { return string(p) }

const (
	storageProviderUnrecognized storageProvider = "unrecognized"
	storageProviderS3           storageProvider = "s3"
	storageProviderGCS          storageProvider = "gs"
	storageProviderHTTP         storageProvider = "http"
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

func storageBugfix(provider string, sa *hy.StorageAccess) *hy.StorageAccess {
	if provider == "gs" {
		// Hybrik has a bug where they identify multi-region GCS -> region GCP
		// transfers as triggering egress costs, so we remove their validation for
		// GCS sources
		sa.MaxCrossRegionMB = -1
	}
	return sa
}

func (p *driver) access(f *job.File, creds string) *hy.StorageAccess {
	if creds == "" {
		if f.Provider() != "gs" {
			return nil
		}
		creds = p.config.GCPCredentialsKey
	}
	return storageBugfix(f.Provider(), &hy.StorageAccess{CredentialsKey: creds})
}

func (p *driver) location(f job.File, creds string) hy.TranscodeLocation {
	return hy.TranscodeLocation{
		StorageProvider: f.Provider(),
		Path:            f.Dir(),
		Access:          p.access(&f, creds),
	}
}

func (p *driver) assetURL(f *job.File, creds string) hy.AssetURL {
	return hy.AssetURL{
		StorageProvider: f.Provider(),
		URL:             f.Name,
		Access:          p.access(f, creds),
	}
}

func (p *driver) asset(f *job.File, creds string, content ...hy.AssetContents) hy.AssetPayload {
	return hy.AssetPayload{
		StorageProvider: f.Provider(),
		URL:             f.Name,
		Contents:        content,
		Access:          p.access(f, creds),
	}
}
