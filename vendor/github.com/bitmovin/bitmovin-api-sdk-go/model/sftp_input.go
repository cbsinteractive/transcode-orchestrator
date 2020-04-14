package model
import (
	"time"
)

type SftpInput struct {
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
	// Host Url or IP of the SFTP server (required)
	Host string `json:"host,omitempty"`
	// Port to use, standard for SFTP: 22
	Port *int32 `json:"port,omitempty"`
	// Use passive mode. Default is true.
	Passive *bool `json:"passive,omitempty"`
	// Your SFTP Username
	Username string `json:"username,omitempty"`
	// Your SFTP password
	Password string `json:"password,omitempty"`
}
func (o SftpInput) InputType() InputType {
    return InputType_SFTP
}

