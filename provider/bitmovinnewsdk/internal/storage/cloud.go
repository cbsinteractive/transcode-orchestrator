package storage

import (
	"net/url"
	"strings"
)

const (
	schemeS3  = "s3"
	schemeGCS = "gs"
)

func parseS3URL(u *url.URL) (bucketName string, objectKey string) {
	return u.Host, strings.TrimLeft(u.Path, "/")
}

func parseGCSURL(u *url.URL) (bucketName string, objectKey string) {
	return u.Host, strings.TrimLeft(u.Path, "/")
}
