package model

type UpdateOrganizationRequest struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// Hexadecimal color
	LabelColor string `json:"labelColor,omitempty"`
}

