package model
// FragmentedMp4MuxingManifestType : Manifest type to be generated for the Fragmented MP4 Muxing. Only significant if a valid fragmentDuration is passed. If this is not set, Smooth Streaming is assumed.
type FragmentedMp4MuxingManifestType string

// List of FragmentedMp4MuxingManifestType
const (
	FragmentedMp4MuxingManifestType_SMOOTH FragmentedMp4MuxingManifestType = "SMOOTH"
	FragmentedMp4MuxingManifestType_DASH_ON_DEMAND FragmentedMp4MuxingManifestType = "DASH_ON_DEMAND"
	FragmentedMp4MuxingManifestType_HLS_BYTE_RANGES FragmentedMp4MuxingManifestType = "HLS_BYTE_RANGES"
	FragmentedMp4MuxingManifestType_NONE FragmentedMp4MuxingManifestType = "NONE"
	FragmentedMp4MuxingManifestType_HLS_BYTE_RANGES_AND_IFRAME_PLAYLIST FragmentedMp4MuxingManifestType = "HLS_BYTE_RANGES_AND_IFRAME_PLAYLIST"
)