package storage

import "github.com/bitmovin/bitmovin-api-sdk-go/model"

// InputAPI holds a collection of media input APIs
type InputAPI struct {
	S3    S3InputAPI
	GCS   GCSInputAPI
	HTTP  HTTPInputAPI
	HTTPS HTTPSInputAPI
}

// S3InputAPI manages media inputs with AWS S3
type S3InputAPI interface {
	Create(model.S3Input) (*model.S3Input, error)
}

// GCSInputAPI manages media inputs with Google Cloud Storage
type GCSInputAPI interface {
	Create(model.GcsInput) (*model.GcsInput, error)
}

// HTTPInputAPI manages media inputs from HTTP sources
type HTTPInputAPI interface {
	Create(model.HttpInput) (*model.HttpInput, error)
}

// HTTPSInputAPI manages media inputs from HTTPS sources
type HTTPSInputAPI interface {
	Create(model.HttpsInput) (*model.HttpsInput, error)
}

// OutputAPI holds a collection of media output APIs
type OutputAPI struct {
	S3  S3OutputAPI
	GCS GCSOutputAPI
}

// S3OutputAPI manages media outputs with AWS S3
type S3OutputAPI interface {
	Create(model.S3Output) (*model.S3Output, error)
}

// GCSOutputAPI manages media outputs with Google Cloud Storage
type GCSOutputAPI interface {
	Create(model.GcsOutput) (*model.GcsOutput, error)
}
