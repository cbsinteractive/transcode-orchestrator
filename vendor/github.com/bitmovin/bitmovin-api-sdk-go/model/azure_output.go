package model
import (
	"time"
)

type AzureOutput struct {
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
	// Azure Account Name (required)
	AccountName string `json:"accountName,omitempty"`
	// Azure Account Key (required)
	AccountKey string `json:"accountKey,omitempty"`
	// Name of the bucket (required)
	Container string `json:"container,omitempty"`
}
func (o AzureOutput) OutputType() OutputType {
    return OutputType_AZURE
}

