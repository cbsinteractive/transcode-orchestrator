package storage

import (
	"fmt"
	"net/url"
)

type pathParser func(*url.URL) string

var pathParsers = map[string]pathParser{
	schemeS3:    s3PathParser,
	schemeGCS:   gcsPathParser,
	schemeHTTP:  httpPathParser,
	schemeHTTPS: httpPathParser,
}

func PathFrom(src string) (string, error) {
	mediaURL, err := url.Parse(src)
	if err != nil {
		return "", fmt.Errorf("could not parse source location %q: %w", src, err)
	}

	parser, found := pathParsers[mediaURL.Scheme]
	if !found {
		return "", fmt.Errorf("invalid scheme %q, only s3, gcs, http, and https urls are supported", mediaURL.Scheme)
	}

	return parser(mediaURL), nil
}

func s3PathParser(src *url.URL) string {
	_, path := parseS3URL(src)
	return path
}

func gcsPathParser(src *url.URL) string {
	_, path := parseGCSURL(src)
	return path
}

func httpPathParser(src *url.URL) string {
	return src.Path
}
