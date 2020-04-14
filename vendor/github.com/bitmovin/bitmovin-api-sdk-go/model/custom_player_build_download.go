package model
import (
	"time"
)

// Custom player download information
type CustomPlayerBuildDownload struct {
	// The link to download the custom built player (required)
	DownloadLink string `json:"downloadLink,omitempty"`
	// Until this date the download link is valid and can be downloaded. (required)
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

