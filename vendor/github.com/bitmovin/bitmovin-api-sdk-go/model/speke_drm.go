package model
import (
	"time"
)

type SpekeDrm struct {
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
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Unique Identifier of the content, will be generated if not set
	ContentId string `json:"contentId,omitempty"`
	// Optional key identifier, will be generated if not set. For SPEKE DRM Configurations with the same contentId and kid the key provider will provide the same keys. 
	Kid string `json:"kid,omitempty"`
	// 16 byte initialization vector represented by a 32-character text string. It is mandatory if systemIds contains AES128 and FairPlay. 
	Iv string `json:"iv,omitempty"`
	// Key provider configuration for SPEKE (required)
	Provider *SpekeDrmProvider `json:"provider,omitempty"`
	// DRM system identifier of the content protection scheme. At minimum one is expected. Not all systemIds are currently supported, support depends on the muxing type.  Relates to SPEKE implementation. See https://dashif.org/identifiers/content_protection/ (required)
	SystemIds []string `json:"systemIds,omitempty"`
}
func (o SpekeDrm) DrmType() DrmType {
    return DrmType_SPEKE
}

