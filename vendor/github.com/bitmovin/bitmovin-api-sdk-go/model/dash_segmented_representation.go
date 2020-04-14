package model

type DashSegmentedRepresentation struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId string `json:"encodingId,omitempty"`
	// UUID of a muxing (required)
	MuxingId string `json:"muxingId,omitempty"`
	// Used to signal a dependency with another representation. The representation may belong to a different adaptation set
	DependencyId string `json:"dependencyId,omitempty"`
	Type DashRepresentationType `json:"type,omitempty"`
	Mode DashRepresentationTypeMode `json:"mode,omitempty"`
	// Path to segments. Will be used as the representation id if the type is set to TEMPLATE_ADAPTATION_SET (required)
	SegmentPath string `json:"segmentPath,omitempty"`
	// Number of the first segment
	StartSegmentNumber *int64 `json:"startSegmentNumber,omitempty"`
	// Number of the last segment. Default is the last one that was encoded
	EndSegmentNumber *int64 `json:"endSegmentNumber,omitempty"`
	// Id of the Keyframe to start with
	StartKeyframeId string `json:"startKeyframeId,omitempty"`
	// Id of the Keyframe to end with
	EndKeyframeId string `json:"endKeyframeId,omitempty"`
}

