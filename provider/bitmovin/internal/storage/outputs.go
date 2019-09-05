package storage

import (
	"net/url"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/config"
	"github.com/pkg/errors"
)

const (
	// we're defaulting all outputs to private
	defaultOutputACL = model.AclPermission_PRIVATE
)

type outputCreator func(*url.URL, OutputAPI, *config.Bitmovin) (outputID string, path string, err error)

var outputCreators = map[string]outputCreator{
	schemeS3:  s3Output,
	schemeGCS: gcsOutput,
}

// NewOutput creates an output and returns an outputId and the folder path or an error
func NewOutput(destLoc string, api OutputAPI, cfg *config.Bitmovin) (outputID string, path string, err error) {
	mediaURL, err := url.Parse(destLoc)
	if err != nil {
		return "", "", errors.Errorf("could not parse destination media location %q: %v", destLoc, err)
	}

	creator, found := outputCreators[mediaURL.Scheme]
	if !found {
		return "", "", errors.Errorf("invalid scheme %q, only s3 and gcs outputs are supported", mediaURL.Scheme)
	}

	return creator(mediaURL, api, cfg)
}

// EncodingOutputFrom returns an encoding output from an output ID and path
func EncodingOutputFrom(outputID, path string) model.EncodingOutput {
	return model.EncodingOutput{
		OutputId:   outputID,
		OutputPath: path,
		Acl:        []model.AclEntry{{Permission: defaultOutputACL}},
	}
}

func s3Output(destURL *url.URL, api OutputAPI, cfg *config.Bitmovin) (string, string, error) {
	bucket, folderPath := parseS3URL(destURL)

	output, err := api.S3.Create(model.S3Output{
		BucketName:  bucket,
		AccessKey:   cfg.AccessKeyID,
		SecretKey:   cfg.SecretAccessKey,
		CloudRegion: model.AwsCloudRegion(cfg.AWSStorageRegion),
		Acl:         []model.AclEntry{{Permission: defaultOutputACL}},
	})
	if err != nil {
		return "", "", errors.Wrap(err, "creating s3 output")
	}

	return output.Id, folderPath, nil
}

func gcsOutput(destURL *url.URL, api OutputAPI, cfg *config.Bitmovin) (string, string, error) {
	bucket, folderPath := parseGCSURL(destURL)

	output, err := api.GCS.Create(model.GcsOutput{
		BucketName:  bucket,
		AccessKey:   cfg.GCSAccessKeyID,
		SecretKey:   cfg.GCSSecretAccessKey,
		CloudRegion: model.GoogleCloudRegion(cfg.GCSStorageRegion),
		Acl:         []model.AclEntry{{Permission: defaultOutputACL}},
	})
	if err != nil {
		return "", "", errors.Wrap(err, "creating gcs output")
	}

	return output.Id, folderPath, nil
}
