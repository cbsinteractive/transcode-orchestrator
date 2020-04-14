package model

type StartLiveEncodingRequest struct {
	// Key for the stream. (a-zA-Z, 3-20 characters) (required)
	StreamKey string `json:"streamKey,omitempty"`
	// List of Hls manifests to use for this live encoding
	HlsManifests []LiveHlsManifest `json:"hlsManifests,omitempty"`
	// List of Dash manifests to use for this live encoding
	DashManifests []LiveDashManifest `json:"dashManifests,omitempty"`
	// The pass mode of the encoding
	LiveEncodingMode EncodingMode `json:"liveEncodingMode,omitempty"`
	// Reupload specific files during a live encoding. This can be helpful if an automatic life cycle policy is enabled on the output storage
	ReuploadSettings *ReuploadSettings `json:"reuploadSettings,omitempty"`
	// Configuration for auto restarting the live encoding
	AutoRestartConfiguration *AutoRestartConfiguration `json:"autoRestartConfiguration,omitempty"`
}

