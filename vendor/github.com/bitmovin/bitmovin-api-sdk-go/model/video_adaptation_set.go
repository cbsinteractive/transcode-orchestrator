package model

type VideoAdaptationSet struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Custom adaptation set attributes
	CustomAttributes []CustomAttribute `json:"customAttributes,omitempty"`
	// Roles of the adaptation set
	Roles []AdaptationSetRole `json:"roles,omitempty"`
	// Provide signaling of CEA 607 and CEA 708
	Accessibilities []Accessibility `json:"accessibilities,omitempty"`
}

