package model

type ReuploadSettings struct {
	// Interval in seconds to reupload the DASH manifest. Valid values are either `null` to never reupload the dash manifest or at least `30`.
	DashManifestInterval *float64 `json:"dashManifestInterval,omitempty"`
	// Interval in seconds to reupload the HLS master file. Valid values are either `0` to never reupload the hls manifest or at least `30`. This is currently not used, as the master file will always be uploaded when one of the playlist files has changed.
	HlsManifestInterval *float64 `json:"hlsManifestInterval,omitempty"`
	// The interval in seconds to reupload the init file for segmented muxings, e.g. fMP4, WebM. Valid values are either `null` to never reupload the init file for segmented muxings or at least `30`.
	MuxingInitFileInterval *float64 `json:"muxingInitFileInterval,omitempty"`
}

