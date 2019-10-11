package hybrik

import (
	"fmt"
	"net/url"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
)

const (
	storageSchemeGCS   = "gs"
	storageSchemeS3    = "s3"
	storageSchemeHTTPS = "https"
	storageSchemeHTTP  = "http"
)

type storageLocation struct {
	provider storageProvider
	path     string
}

func (p *hybrikProvider) transcodeLocationFrom(dest storageLocation, env db.ExecutionEnvironment) hybrik.TranscodeLocation {
	location := hybrik.TranscodeLocation{
		StorageProvider: dest.provider.string(),
		Path:            dest.path,
	}

	if access, add := p.storageAccessFrom(dest.provider, env); add {
		location.Access = access
	}

	return location
}

func (p *hybrikProvider) assetURLFrom(dest storageLocation, env db.ExecutionEnvironment) hybrik.AssetURL {
	assetURL := hybrik.AssetURL{
		StorageProvider: dest.provider.string(),
		URL:             dest.path,
	}

	if access, add := p.storageAccessFrom(dest.provider, env); add {
		assetURL.Access = access
	}

	return assetURL
}

func (p *hybrikProvider) assetPayloadFrom(provider storageProvider, url string, contents []hybrik.AssetContents, env db.ExecutionEnvironment) hybrik.AssetPayload {
	assetPayload := hybrik.AssetPayload{
		StorageProvider: provider.string(),
		URL:             url,
		Contents:        contents,
	}

	if access, add := p.storageAccessFrom(provider, env); add {
		assetPayload.Access = access
	}

	return assetPayload
}

func (p *hybrikProvider) storageAccessFrom(provider storageProvider, env db.ExecutionEnvironment) (*hybrik.StorageAccess, bool) {
	var maxCrossRegionMB int

	// Hybrik has a bug where they identify multi-region GCS -> region GCP
	// transfers as triggering egress costs, so we remove their validation for
	// GCS sources
	if provider == storageProviderGCS {
		maxCrossRegionMB = -1
	}

	if alias := env.CredentialsAlias; alias != "" {
		return &hybrik.StorageAccess{CredentialsKey: alias, MaxCrossRegionMB: maxCrossRegionMB}, true
	}

	if provider == storageProviderGCS {
		return &hybrik.StorageAccess{CredentialsKey: p.config.GCPCredentialsKey, MaxCrossRegionMB: maxCrossRegionMB}, true
	}

	return nil, false
}

func storageProviderFrom(path string) (storageProvider, error) {
	u, err := url.Parse(path)
	if err != nil {
		return storageProviderUnrecognized, err
	}

	switch u.Scheme {
	case storageSchemeS3:
		return storageProviderS3, nil
	case storageSchemeGCS:
		return storageProviderGCS, nil
	case storageSchemeHTTPS:
		return storageProviderHTTP, nil
	case storageSchemeHTTP:
		return storageProviderHTTP, nil
	}

	return storageProviderUnrecognized, fmt.Errorf("the scheme %q is unsupported", u.Scheme)
}
