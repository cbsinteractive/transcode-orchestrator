package storage

import (
	"testing"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/config"
	"github.com/cbsinteractive/video-transcoding-api/test"
	"github.com/pkg/errors"
)

func TestNewInput(t *testing.T) {
	tests := []struct {
		name, srcMedia            string
		api                       InputAPI
		cfg                       config.Bitmovin
		wantID, wantPath, wantErr string
		assertInputParams         func(*testing.T, InputAPI)
	}{
		{
			name:     "a valid aws config is properly parsed",
			srcMedia: "s3://bucket_name/path/to/file.mp4",
			api:      fakeInputAPIReturningInputID("some-input-id"),
			cfg: config.Bitmovin{
				AccessKeyID:      "access-key",
				SecretAccessKey:  "secret-key",
				AWSStorageRegion: "us-west-2",
			},
			wantID:   "some-input-id",
			wantPath: "path/to/file.mp4",
			assertInputParams: func(t *testing.T, api InputAPI) {
				fakeS3API := api.S3.(*fakeS3InputAPI)

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
			name:     "a valid gcs config is properly parsed",
			srcMedia: "gs://gcs_bucket_name/path/to/gcs_file.mp4",
			api:      fakeInputAPIReturningInputID("some-gcs-input-id"),
			cfg: config.Bitmovin{
				GCSAccessKeyID:     "gcs-access-key",
				GCSSecretAccessKey: "gcs-secret-key",
				GCSStorageRegion:   "us-east1",
			},
			wantID:   "some-gcs-input-id",
			wantPath: "path/to/gcs_file.mp4",
			assertInputParams: func(t *testing.T, api InputAPI) {
				fakeGCSAPI := api.GCS.(*fakeGCSInputAPI)

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
			name:     "a valid http config is properly parsed",
			srcMedia: "http://somedomain.com/path/to/http_file.mp4",
			api:      fakeInputAPIReturningInputID("some-http-input-id"),
			wantID:   "some-http-input-id",
			wantPath: "/path/to/http_file.mp4",
			assertInputParams: func(t *testing.T, api InputAPI) {
				fakeHTTPAPI := api.HTTP.(*fakeHTTPInputAPI)

				if g, e := fakeHTTPAPI.createdWithHost, "somedomain.com"; g != e {
					t.Errorf("invalid host name: got %q, expected %q", g, e)
				}
			},
		},
		{
			name:     "a valid https config is properly parsed",
			srcMedia: "https://somehttpsdomain.com/path/to/https_file.mp4",
			api:      fakeInputAPIReturningInputID("some-https-input-id"),
			wantID:   "some-https-input-id",
			wantPath: "/path/to/https_file.mp4",
			assertInputParams: func(t *testing.T, api InputAPI) {
				fakeHTTPSAPI := api.HTTPS.(*fakeHTTPSInputAPI)

				if g, e := fakeHTTPSAPI.createdWithHost, "somehttpsdomain.com"; g != e {
					t.Errorf("invalid host name: got %q, expected %q", g, e)
				}
			},
		},
		{
			name:     "an unsupported src url results in an error",
			srcMedia: "cbscloud://some-bucket/some/path/file.mp4",
			api:      fakeInputAPIReturningInputID("some-https-input-id"),
			wantErr:  `invalid scheme "cbscloud", only s3, gcs, http, and https urls are supported`,
		},
		{
			name:     "if the bitmovin api is erroring, we get a useful error when creating s3 inputs",
			srcMedia: "s3://bucket/path/to/file.mp4",
			api:      erroringInputAPI,
			wantErr:  "creating s3 input: fake error from api",
		},
		{
			name:     "if the bitmovin api is erroring, we get a useful error when creating gcs inputs",
			srcMedia: "gs://bucket/path/to/file.mp4",
			api:      erroringInputAPI,
			wantErr:  "creating gcs input: fake error from api",
		},

		{
			name:     "if the bitmovin api is erroring, we get a useful error when creating http inputs",
			srcMedia: "http://somehttpdomain.com/path/to/http_file.mp4",
			api:      erroringInputAPI,
			wantErr:  "creating http input: fake error from api",
		},

		{
			name:     "if the bitmovin api is erroring, we get a useful error when creating https inputs",
			srcMedia: "https://somehttpsdomain.com/path/to/https_file.mp4",
			api:      erroringInputAPI,
			wantErr:  "creating https input: fake error from api",
		},
		{
			name:     "an unparsable src url results in a useful error",
			srcMedia: "s3://%%some-bucket/some/path/file.mp4",
			api:      fakeInputAPIReturningInputID("some-input-id"),
			wantErr:  `could not parse source media location "s3://%%some-bucket/some/path/file.mp4"`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			id, path, err := NewInput(tt.srcMedia, tt.api, &tt.cfg)
			if shouldReturn := test.AssertWantErr(err, tt.wantErr, "NewInput()", t); shouldReturn {
				return
			}

			if g, e := id, tt.wantID; g != e {
				t.Errorf("invalid id returned, got %q, expected %q", g, e)
			}

			if g, e := path, tt.wantPath; g != e {
				t.Errorf("invalid path returned, got %q, expected %q", g, e)
			}

			if inputAssertion := tt.assertInputParams; inputAssertion != nil {
				inputAssertion(t, tt.api)
			}
		})
	}
}

func fakeInputAPIReturningInputID(id string) InputAPI {
	return InputAPI{
		S3:    &fakeS3InputAPI{inputIDToReturn: id},
		GCS:   &fakeGCSInputAPI{inputIDToReturn: id},
		HTTP:  &fakeHTTPInputAPI{inputIDToReturn: id},
		HTTPS: &fakeHTTPSInputAPI{inputIDToReturn: id},
	}
}

type fakeGCSInputAPI struct {
	inputIDToReturn            string
	createdWithBucket          string
	createdWithAccessKeyID     string
	createdWithSecretAccessKey string
	createdWithRegion          model.GoogleCloudRegion
}

func (a *fakeGCSInputAPI) Create(input model.GcsInput) (*model.GcsInput, error) {
	a.createdWithBucket = input.BucketName
	a.createdWithAccessKeyID = input.AccessKey
	a.createdWithSecretAccessKey = input.SecretKey
	a.createdWithRegion = input.CloudRegion
	return &model.GcsInput{Id: a.inputIDToReturn}, nil
}

type fakeS3InputAPI struct {
	inputIDToReturn            string
	createdWithBucket          string
	createdWithAccessKeyID     string
	createdWithSecretAccessKey string
	createdWithRegion          model.AwsCloudRegion
}

func (a *fakeS3InputAPI) Create(input model.S3Input) (*model.S3Input, error) {
	a.createdWithBucket = input.BucketName
	a.createdWithAccessKeyID = input.AccessKey
	a.createdWithSecretAccessKey = input.SecretKey
	a.createdWithRegion = input.CloudRegion
	return &model.S3Input{Id: a.inputIDToReturn}, nil
}

type fakeHTTPInputAPI struct {
	inputIDToReturn string
	createdWithHost string
}

func (a *fakeHTTPInputAPI) Create(input model.HttpInput) (*model.HttpInput, error) {
	a.createdWithHost = input.Host
	return &model.HttpInput{Id: a.inputIDToReturn}, nil
}

type fakeHTTPSInputAPI struct {
	inputIDToReturn string
	createdWithHost string
}

func (a *fakeHTTPSInputAPI) Create(input model.HttpsInput) (*model.HttpsInput, error) {
	a.createdWithHost = input.Host
	return &model.HttpsInput{Id: a.inputIDToReturn}, nil
}

var erroringInputAPI = InputAPI{
	S3:    &erroringS3InputAPI{},
	GCS:   &erroringGCSInputAPI{},
	HTTP:  &erroringHTTPInputAPI{},
	HTTPS: &erroringHTTPSInputAPI{},
}

type erroringGCSInputAPI struct{}

func (*erroringGCSInputAPI) Create(model.GcsInput) (*model.GcsInput, error) {
	return nil, errors.New("fake error from api")
}

type erroringS3InputAPI struct{}

func (*erroringS3InputAPI) Create(model.S3Input) (*model.S3Input, error) {
	return nil, errors.New("fake error from api")
}

type erroringHTTPInputAPI struct{}

func (*erroringHTTPInputAPI) Create(model.HttpInput) (*model.HttpInput, error) {
	return nil, errors.New("fake error from api")
}

type erroringHTTPSInputAPI struct{}

func (*erroringHTTPSInputAPI) Create(model.HttpsInput) (*model.HttpsInput, error) {
	return nil, errors.New("fake error from api")
}
