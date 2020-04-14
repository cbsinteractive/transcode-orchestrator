package model
import (
	"time"
)

type AwsAccount struct {
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
	// Amazon access key (required)
	AccessKey string `json:"accessKey,omitempty"`
	// Amazon secret key (required)
	SecretKey string `json:"secretKey,omitempty"`
	// Amazon account number (12 digits as per AWS spec) (required)
	AccountNumber string `json:"accountNumber,omitempty"`
}

