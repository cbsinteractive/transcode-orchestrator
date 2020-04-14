package model
import (
	"time"
)

type FtpInput struct {
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
	// Host URL or IP of the FTP server (required)
	Host string `json:"host,omitempty"`
	// Port to use, standard for FTP: 21
	Port *int32 `json:"port,omitempty"`
	// Use passive mode. Default is true.
	Passive *bool `json:"passive,omitempty"`
	// Your FTP Username
	Username string `json:"username,omitempty"`
	// Your FTP password
	Password string `json:"password,omitempty"`
	// Ensure that connections originate from the declared ftp host. Default is true.
	RemoteVerificationEnabled *bool `json:"remoteVerificationEnabled,omitempty"`
}
func (o FtpInput) InputType() InputType {
    return InputType_FTP
}

