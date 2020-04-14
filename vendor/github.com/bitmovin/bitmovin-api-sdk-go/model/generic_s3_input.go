package model
import (
	"time"
)

type GenericS3Input struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Your generic S3 bucket name (required)
	BucketName string `json:"bucketName,omitempty"`
	// The generic S3 server hostname (or IP address) (required)
	Host string `json:"host,omitempty"`
	// The port on which the generic S3 server is running on (if not provided 8000 will be used)
	Port *int32 `json:"port,omitempty"`
	// Controls whether SSL is used or not
	Ssl *bool `json:"ssl,omitempty"`
	// Specifies the method used for authentication
	SignatureVersion S3SignatureVersion `json:"signatureVersion,omitempty"`
	// Your generic S3 access key (required)
	AccessKey string `json:"accessKey,omitempty"`
	// Your generic S3 secret key (required)
	SecretKey string `json:"secretKey,omitempty"`
}
func (o GenericS3Input) InputType() InputType {
    return InputType_GENERIC_S3
}

