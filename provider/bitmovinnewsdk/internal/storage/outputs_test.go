package storage

import (
	"testing"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/config"
	"github.com/cbsinteractive/video-transcoding-api/test"
	"github.com/pkg/errors"
)

func TestNewOutput(t *testing.T) {
	tests := []struct {
		name, destLoc             string
		api                       OutputAPI
		cfg                       config.Bitmovin
		wantID, wantPath, wantErr string
		assertOutputParams        func(*testing.T, OutputAPI)
	}{
		{
			name:    "a valid aws config is properly parsed",
			destLoc: "s3://bucket_name/path/to/file.mp4",
			api:     fakeOutputAPIReturningOutputID("some-output-id"),
			cfg: config.Bitmovin{
				AccessKeyID:      "access-key",
				SecretAccessKey:  "secret-key",
				AWSStorageRegion: "us-west-2",
			},
			wantID:   "some-output-id",
			wantPath: "path/to/file.mp4",
			assertOutputParams: func(t *testing.T, api OutputAPI) {
				fakeS3API := api.S3.(*fakeS3OutputAPI)

				if g, e := fakeS3API.createdWithBucket, "bucket_name"; g != e {
					t.Errorf("invalid bucket name: got %q, expected %q", g, e)
				}

				if g, e := fakeS3API.createdWithAccessKeyID, "access-key"; g != e {
					t.Errorf("invalid access key: got %q, expected %q", g, e)
				}

				if g, e := fakeS3API.createdWithSecretAccessKey, "secret-key"; g != e {
					t.Errorf("invalid secret access key: got %q, expected %q", g, e)
				}

				if g, e := fakeS3API.createdWithRegion, model.AwsCloudRegion("us-west-2"); g != e {
					t.Errorf("invalid region: got %q, expected %q", g, e)
				}
			},
		},
		{
			name:    "a valid gcs config is properly parsed",
			destLoc: "gs://gcs_bucket_name/path/to/gcs_file.mp4",
			api:     fakeOutputAPIReturningOutputID("some-gcs-output-id"),
			cfg: config.Bitmovin{
				GCSAccessKeyID:     "gcs-access-key",
				GCSSecretAccessKey: "gcs-secret-key",
				GCSStorageRegion:   "us-east1",
			},
			wantID:   "some-gcs-output-id",
			wantPath: "path/to/gcs_file.mp4",
			assertOutputParams: func(t *testing.T, api OutputAPI) {
				fakeGCSAPI := api.GCS.(*fakeGCSOutputAPI)

				if g, e := fakeGCSAPI.createdWithBucket, "gcs_bucket_name"; g != e {
					t.Errorf("invalid bucket name: got %q, expected %q", g, e)
				}

				if g, e := fakeGCSAPI.createdWithAccessKeyID, "gcs-access-key"; g != e {
					t.Errorf("invalid access key: got %q, expected %q", g, e)
				}

				if g, e := fakeGCSAPI.createdWithSecretAccessKey, "gcs-secret-key"; g != e {
					t.Errorf("invalid secret access key: got %q, expected %q", g, e)
				}

				if g, e := fakeGCSAPI.createdWithRegion, model.GoogleCloudRegion("us-east1"); g != e {
					t.Errorf("invalid region: got %q, expected %q", g, e)
				}
			},
		},
		{
			name:    "an unsupported src url results in an error",
			destLoc: "cbscloud://some-bucket/some/path/file.mp4",
			api:     fakeOutputAPIReturningOutputID("some-https-output-id"),
			wantErr: `invalid scheme "cbscloud", only s3 and gcs outputs are supported`,
		},
		{
			name:    "if the bitmovin api is erroring, we get a useful error when creating s3 outputs",
			destLoc: "s3://bucket/path/to/file.mp4",
			api:     erroringOutputAPI,
			wantErr: "creating s3 output: fake error from api",
		},
		{
			name:    "if the bitmovin api is erroring, we get a useful error when creating gcs outputs",
			destLoc: "gs://bucket/path/to/file.mp4",
			api:     erroringOutputAPI,
			wantErr: "creating gcs output: fake error from api",
		},
		{
			name:    "an unparsable src url results in a useful error",
			destLoc: "s3://%%some-bucket/some/path/file.mp4",
			api:     fakeOutputAPIReturningOutputID("some-output-id"),
			wantErr: `could not parse destination media location "s3://%%some-bucket/some/path/file.mp4": ` +
				`parse s3://%%some-bucket/some/path/file.mp4: invalid URL escape "%%s"`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			id, path, err := NewOutput(tt.destLoc, tt.api, &tt.cfg)
			if shouldReturn := test.AssertWantErr(err, tt.wantErr, "NewOutput()", t); shouldReturn {
				return
			}

			if g, e := id, tt.wantID; g != e {
				t.Errorf("invalid id returned, got %q, expected %q", g, e)
			}

			if g, e := path, tt.wantPath; g != e {
				t.Errorf("invalid path returned, got %q, expected %q", g, e)
			}

			if outputAssertion := tt.assertOutputParams; outputAssertion != nil {
				outputAssertion(t, tt.api)
			}
		})
	}
}

func fakeOutputAPIReturningOutputID(id string) OutputAPI {
	return OutputAPI{
		GCS: &fakeGCSOutputAPI{outputIDToReturn: id},
		S3:  &fakeS3OutputAPI{outputIDToReturn: id},
	}
}

type fakeGCSOutputAPI struct {
	outputIDToReturn           string
	createdWithBucket          string
	createdWithAccessKeyID     string
	createdWithSecretAccessKey string
	createdWithRegion          model.GoogleCloudRegion
}

func (a *fakeGCSOutputAPI) Create(output model.GcsOutput) (*model.GcsOutput, error) {
	a.createdWithBucket = output.BucketName
	a.createdWithAccessKeyID = output.AccessKey
	a.createdWithSecretAccessKey = output.SecretKey
	a.createdWithRegion = output.CloudRegion
	return &model.GcsOutput{Id: a.outputIDToReturn}, nil
}

type fakeS3OutputAPI struct {
	outputIDToReturn           string
	createdWithBucket          string
	createdWithAccessKeyID     string
	createdWithSecretAccessKey string
	createdWithRegion          model.AwsCloudRegion
}

func (a *fakeS3OutputAPI) Create(output model.S3Output) (*model.S3Output, error) {
	a.createdWithBucket = output.BucketName
	a.createdWithAccessKeyID = output.AccessKey
	a.createdWithSecretAccessKey = output.SecretKey
	a.createdWithRegion = output.CloudRegion
	return &model.S3Output{Id: a.outputIDToReturn}, nil
}

var erroringOutputAPI = OutputAPI{
	GCS: &erroringGCSOutputAPI{},
	S3:  &erroringS3OutputAPI{},
}

type erroringGCSOutputAPI struct{}

func (*erroringGCSOutputAPI) Create(model.GcsOutput) (*model.GcsOutput, error) {
	return nil, errors.New("fake error from api")
}

type erroringS3OutputAPI struct{}

func (*erroringS3OutputAPI) Create(model.S3Output) (*model.S3Output, error) {
	return nil, errors.New("fake error from api")
}
