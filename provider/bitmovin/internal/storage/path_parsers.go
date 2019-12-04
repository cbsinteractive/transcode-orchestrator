package storage

import (
	"fmt"
	"net/url"
)

func PathFrom(src string) (string, error) {
	mediaURL, err := url.Parse(src)
	if err != nil {
		return "", fmt.Errorf("could not parse source location %q: %w", src, err)
	}

	return mediaURL.RequestURI(), nil
}
