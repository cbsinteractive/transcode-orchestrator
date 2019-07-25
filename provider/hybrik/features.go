package hybrik

// SegmentedRendering holds segmented rendering parameters
type SegmentedRendering struct {
	// Duration (in seconds) of a segment in segment encode mode. Minimum: 1
	Duration int `json:"duration,omitempty"`

	// Duration (in seconds) to look for a dominant previous or following scene change. Note that
	// the segment duration can then be up to Duration + SceneChangeSearchDuration long.
	SceneChangeSearchDuration int `json:"sceneChangeSearchDuration,omitempty"`

	// Total number of segments
	NumTotalSegments int `json:"totalSegments,omitempty"`

	// Combiner will merge and re-stripe transport streams
	EnableStrictCFR bool `json:"strictCFR,omitempty"`

	// Timebase offset to be used by the muxer
	MuxTimebaseOffset int `json:"muxOffsetOTB,omitempty"`
}

// DolbyVision hold configuration options for the DEE
type DolbyVision struct {
	InputEDRAspect  string `json:"inputEDRAspect"`
	InputEDRPad     string `json:"inputEDRPad"`
	InputEDRCrop    string `json:"inputEDRCrop"`
	VesQCVersion    string `json:"vesQCVersion"`
	MezzQCVersion   string `json:"mezzQCVersion"`
	Mp4QCVersion    string `json:"mp4QCVersion"`
	Mp4MuxerVersion string `json:"mp4MuxerVersion"`
	SDKVersion      string `json:"sdkVersion"`
}
