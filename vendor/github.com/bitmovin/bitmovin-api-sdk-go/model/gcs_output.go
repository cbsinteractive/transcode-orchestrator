package model
import (
	"time"
)

type GcsOutput struct {
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
	// GCS access key (required)
	AccessKey string `json:"accessKey,omitempty"`
	// GCS secret key (required)
	SecretKey string `json:"secretKey,omitempty"`
	// Name of the bucket (required)
	BucketName string `json:"bucketName,omitempty"`
	CloudRegion GoogleCloudRegion `json:"cloudRegion,omitempty"`
}
func (o GcsOutput) OutputType() OutputType {
    return OutputType_GCS
}

