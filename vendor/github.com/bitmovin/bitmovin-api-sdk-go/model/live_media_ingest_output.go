package model
import (
	"time"
)

// Live Media and Metadata Ingest Protocol. See https://tools.ietf.org/html/draft-mekuria-mmediaingest-01.
type LiveMediaIngestOutput struct {
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
	// URL specifying the publishing point for the output. Can use either http or https. (required)
	PublishingPoint string `json:"publishingPoint,omitempty"`
}
func (o LiveMediaIngestOutput) OutputType() OutputType {
    return OutputType_LIVE_MEDIA_INGEST
}

