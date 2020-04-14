package model

type CustomXmlElement struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// String representation of the XML element (required)
	Data string `json:"data,omitempty"`
}

