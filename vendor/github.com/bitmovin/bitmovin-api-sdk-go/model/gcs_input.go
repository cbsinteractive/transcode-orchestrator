package model
import (
	"time"
)

type GcsInput struct {
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
	// Name of the bucket (required)
	BucketName string `json:"bucketName,omitempty"`
	// The cloud region in which the bucket is located. Is used to determine the ideal location for your encodings automatically.
	CloudRegion GoogleCloudRegion `json:"cloudRegion,omitempty"`
	// GCS access key (required)
	AccessKey string `json:"accessKey,omitempty"`
	// GCS secret key (required)
	SecretKey string `json:"secretKey,omitempty"`
}
func (o GcsInput) InputType() InputType {
    return InputType_GCS
}

