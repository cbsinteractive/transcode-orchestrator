package model
import (
	"time"
)

type AkamaiMslOutput struct {
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
	// The Akamai stream ID (required)
	StreamId *int32 `json:"streamId,omitempty"`
	// The Akamai event name (required)
	EventName string `json:"eventName,omitempty"`
	StreamFormat AkamaiMslStreamFormat `json:"streamFormat,omitempty"`
	MslVersion AkamaiMslVersion `json:"mslVersion,omitempty"`
}
func (o AkamaiMslOutput) OutputType() OutputType {
    return OutputType_AKAMAI_MSL
}

