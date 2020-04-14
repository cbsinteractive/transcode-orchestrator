package model
// DrmType : The type of the drm
type DrmType string

// List of DrmType
const (
	DrmType_WIDEVINE DrmType = "WIDEVINE"
	DrmType_PLAYREADY DrmType = "PLAYREADY"
	DrmType_PRIMETIME DrmType = "PRIMETIME"
	DrmType_FAIRPLAY DrmType = "FAIRPLAY"
	DrmType_MARLIN DrmType = "MARLIN"
	DrmType_CLEARKEY DrmType = "CLEARKEY"
	DrmType_AES DrmType = "AES"
	DrmType_CENC DrmType = "CENC"
	DrmType_SPEKE DrmType = "SPEKE"
)