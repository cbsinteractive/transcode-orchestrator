package model

type UtcTiming struct {
	// The server to get the time from (required)
	Value string `json:"value,omitempty"`
	// The scheme id to use. Please refer to the DASH standard. (required)
	SchemeIdUri string `json:"schemeIdUri,omitempty"`
}

