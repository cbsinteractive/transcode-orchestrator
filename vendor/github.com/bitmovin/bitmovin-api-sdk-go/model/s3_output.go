package model
import (
	"time"
)

type S3Output struct {
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
	Acl []AclEntry `json:"acl,omitempty"`
	// Amazon S3 bucket name (required)
	BucketName string `json:"bucketName,omitempty"`
	// Amazon S3 access key (required)
	AccessKey string `json:"accessKey,omitempty"`
	// Amazon S3 secret key (required)
	SecretKey string `json:"secretKey,omitempty"`
	// If set a user defined tag (x-amz-meta-) with that key will be used to store the MD5 hash of the file.
	Md5MetaTag string `json:"md5MetaTag,omitempty"`
	CloudRegion AwsCloudRegion `json:"cloudRegion,omitempty"`
	SignatureVersion S3SignatureVersion `json:"signatureVersion,omitempty"`
}
func (o S3Output) OutputType() OutputType {
    return OutputType_S3
}

