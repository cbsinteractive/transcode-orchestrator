package model
type ManifestType string

// List of ManifestType
const (
	ManifestType_DASH ManifestType = "DASH"
	ManifestType_HLS ManifestType = "HLS"
	ManifestType_SMOOTH_STREAMING ManifestType = "SMOOTH_STREAMING"
)