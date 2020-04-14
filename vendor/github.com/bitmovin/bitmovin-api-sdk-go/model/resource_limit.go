package model

type ResourceLimit struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	// Specifies the remaining limit value for limits divided to sub-organizations. This property is returned only for parent organizations and only for mentioned limits.
	EffectiveValue string `json:"effectiveValue,omitempty"`
}

