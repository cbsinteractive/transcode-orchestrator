package hybrik

import (
	"fmt"
	"net/url"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
)

const (
	storageSchemeGCS = "gs"
	storageSchemeS3  = "s3"
)

type storageLocation struct {
	provider storageProvider
	path     string
}

func (p *hybrikProvider) transcodeLocationFrom(dest storageLocation, env db.ExecutionEnvironment) hybrik.TranscodeLocation {
	location := hybrik.TranscodeLocation{
		StorageProvider: dest.provider,
		Path:            dest.path,
	}

	if access, add := p.storageAccessFrom(dest.provider, env); add {
		location.Access = access
	}

	return location
}

func (p *hybrikProvider) assetURLFrom(dest storageLocation, env db.ExecutionEnvironment) hybrik.AssetURL {
	assetURL := hybrik.AssetURL{
		StorageProvider: dest.provider,
		URL:             dest.path,
	}

	if access, add := p.storageAccessFrom(dest.provider, env); add {
		assetURL.Access = access
	}

	return assetURL
}

func (p *hybrikProvider) assetPayloadFrom(provider, url string, contents []hybrik.AssetContents, env db.ExecutionEnvironment) hybrik.AssetPayload {
	assetPayload := hybrik.AssetPayload{
		StorageProvider: provider,
		URL:             url,
		Contents:        contents,
	}

	if access, add := p.storageAccessFrom(provider, env); add {
		assetPayload.Access = access
	}

	return assetPayload
}

func (p *hybrikProvider) storageAccessFrom(provider string, env db.ExecutionEnvironment) (*hybrik.StorageAccess, bool) {
	if alias := env.CredentialsAlias; alias != "" {
		return &hybrik.StorageAccess{CredentialsKey: alias}, true
	}

	if provider == storageProviderGCS {
		return &hybrik.StorageAccess{CredentialsKey: p.config.GCPCredentialsKey}, true
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
	}

	return storageProviderUnrecognized, fmt.Errorf("the scheme %q is unsupported", u.Scheme)
}
