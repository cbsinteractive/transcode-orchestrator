package hybrik

import (
	"fmt"
	"net/url"
)

const (
	storageSchemeGCS = "gs"
	storageSchemeS3  = "s3"
)

type storageLocation struct {
	provider storageProvider
	path     string
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
