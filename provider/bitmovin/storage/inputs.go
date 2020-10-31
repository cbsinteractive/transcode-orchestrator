package storage

import (
	"fmt"
	"net/url"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/pkg/errors"
)

const (
	schemeHTTP  = "http"
	schemeHTTPS = "https"
)

type inputCreator func(*url.URL, InputAPI, *config.Bitmovin) (inputID string, err error)

var inputCreators = map[string]inputCreator{
	schemeS3:    s3Input,
	schemeGCS:   gcsInput,
	schemeHTTP:  httpInput,
	schemeHTTPS: httpsInput,
}

// NewInput creates an input and returns an inputID and the media path or an error
func NewInput(srcMediaLoc string, api InputAPI, cfg *config.Bitmovin) (inputID string, err error) {
	mediaURL, err := url.Parse(srcMediaLoc)
	if err != nil {
		return "", fmt.Errorf("could not parse source media location %q", srcMediaLoc)
	}

	creator, found := inputCreators[mediaURL.Scheme]
	if !found {
		return "", fmt.Errorf("invalid scheme %q, only s3, gcs, http, and https urls are supported", mediaURL.Scheme)
	}

	return creator(mediaURL, api, cfg)
}

func s3Input(srcMediaURL *url.URL, api InputAPI, cfg *config.Bitmovin) (inputID string, err error) {
	bucket, _ := parseS3URL(srcMediaURL)

	input, err := api.S3.Create(model.S3Input{
		CloudRegion: model.AwsCloudRegion(cfg.AWSStorageRegion),
		BucketName:  bucket,
		AccessKey:   cfg.AccessKeyID,
		SecretKey:   cfg.SecretAccessKey,
	})
	if err != nil {
		return "", errors.Wrap(err, "creating s3 input")
	}

	return input.Id, nil
}

func gcsInput(srcMediaURL *url.URL, api InputAPI, cfg *config.Bitmovin) (inputID string, err error) {
	bucket, _ := parseGCSURL(srcMediaURL)

	input, err := api.GCS.Create(model.GcsInput{
		CloudRegion: model.GoogleCloudRegion(cfg.GCSStorageRegion),
		BucketName:  bucket,
		AccessKey:   cfg.GCSAccessKeyID,
		SecretKey:   cfg.GCSSecretAccessKey,
	})
	if err != nil {
		return "", errors.Wrap(err, "creating gcs input")
	}

	return input.Id, nil
}

func httpInput(srcMediaURL *url.URL, api InputAPI, _ *config.Bitmovin) (inputID string, err error) {
	input, err := api.HTTP.Create(model.HttpInput{
		Host: srcMediaURL.Host,
	})
	if err != nil {
		return "", errors.Wrap(err, "creating http input")
	}

	return input.Id, nil
}

func httpsInput(srcMediaURL *url.URL, api InputAPI, _ *config.Bitmovin) (inputID string, err error) {
	input, err := api.HTTPS.Create(model.HttpsInput{
		Host: srcMediaURL.Host,
	})
	if err != nil {
		return "", errors.Wrap(err, "creating https input")
	}

	return input.Id, nil
}
