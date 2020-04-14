package model

type ReprioritizeEncodingRequest struct {
	// Priority of the Encoding (required)
	Priority *int32 `json:"priority,omitempty"`
}

