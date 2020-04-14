package model

type ApiErrorDefinition struct {
	// The error code.
	Code *int64 `json:"code,omitempty"`
	// The error category.
	Category string `json:"category,omitempty"`
	// The returned error description.
	Description string `json:"description,omitempty"`
}

