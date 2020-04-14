package model
import (
	"time"
)

type SrtInput struct {
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
	// The SRT mode to use (required)
	Mode SrtMode `json:"mode,omitempty"`
	// The name or IP of the host providing the SRT stream (only used in CALLER mode)
	Host string `json:"host,omitempty"`
	// The port to connect to or listen on. Has to be one of [2088, 2089, 2090, 2091] when using LISTENER mode. (required)
	Port *int32 `json:"port,omitempty"`
	// The path parameter of the SRT stream
	Path string `json:"path,omitempty"`
	// The maximum accepted transmission latency in milliseconds (when both parties set different values, the maximum of the two is used for both)
	Latency *int32 `json:"latency,omitempty"`
	// The passphrase used to secure the SRT stream. For AES-128 encryption, you must enter a 16-character passphrase; for AES-256, you must enter a 32-character passphrase
	Passphrase string `json:"passphrase,omitempty"`
	// The type of AES encryption determines the length of the key (passphrase). AES-128 uses a 16-character (128-bit) passphrase, and AES-256 uses a 32-character (256-bit) passphrase.
	KeyLength *int32 `json:"keyLength,omitempty"`
	BackupSrtInputs *BackupSrtInputs `json:"backupSrtInputs,omitempty"`
}
func (o SrtInput) InputType() InputType {
    return InputType_SRT
}

