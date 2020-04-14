package model

type XmlNamespace struct {
	// Name of the XML Namespace reference (required)
	Prefix string `json:"prefix,omitempty"`
	// Source of the XML Namespace reference (required)
	Uri string `json:"uri,omitempty"`
}

