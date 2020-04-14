package model
import (
	"time"
)

type CencDrm struct {
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
	// 16 byte encryption key, 32 hexadecimal characters (required)
	Key string `json:"key,omitempty"`
	// 16 byte encryption key id. Required for any other DRM but FairPlay
	Kid string `json:"kid,omitempty"`
	// The encryption method to use. Default is `CTR` (required)
	EncryptionMode EncryptionMode `json:"encryptionMode,omitempty"`
	// Size of the initialization vector
	IvSize IvSize `json:"ivSize,omitempty"`
	// Enables compatibility with the Protected Interoperable File Format (PIFF) specification
	EnablePiffCompatibility *bool `json:"enablePiffCompatibility,omitempty"`
	// Configuration for Widevine DRM
	Widevine *CencWidevine `json:"widevine,omitempty"`
	// Configuration for PlayReady DRM
	PlayReady *CencPlayReady `json:"playReady,omitempty"`
	// Configuration for Marlin DRM
	Marlin *CencMarlin `json:"marlin,omitempty"`
	// Configuration for FairPlay DRM
	FairPlay *CencFairPlay `json:"fairPlay,omitempty"`
}
func (o CencDrm) DrmType() DrmType {
    return DrmType_CENC
}

