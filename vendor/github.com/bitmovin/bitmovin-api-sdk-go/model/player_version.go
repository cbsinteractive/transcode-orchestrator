package model
import (
	"time"
)

type PlayerVersion struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Version of the Player (required)
	Version string `json:"version,omitempty"`
	// URL of the specified player (required)
	CdnUrl string `json:"cdnUrl,omitempty"`
	// Download URL of the specified player package (required)
	DownloadUrl string `json:"downloadUrl,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ (required)
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

