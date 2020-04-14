package model

type VttMediaInfo struct {
	// The value is a quoted-string which specifies the group to which the Rendition belongs. (required)
	GroupId string `json:"groupId,omitempty"`
	// Primary language in the rendition.
	Language string `json:"language,omitempty"`
	// Identifies a language that is associated with the Rendition.
	AssocLanguage string `json:"assocLanguage,omitempty"`
	// Human readable description of the rendition. (required)
	Name string `json:"name,omitempty"`
	// If set to true, the client SHOULD play this Rendition of the content in the absence of information from the user.
	IsDefault *bool `json:"isDefault,omitempty"`
	// If set to true, the client MAY choose to play this Rendition in the absence of explicit user preference.
	Autoselect *bool `json:"autoselect,omitempty"`
	// Contains Uniform Type Identifiers
	Characteristics []string `json:"characteristics,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// The URL of the referenced VTT file (required)
	VttUrl string `json:"vttUrl,omitempty"`
	// The URI of the Rendition (required)
	Uri string `json:"uri,omitempty"`
	// A value of true indicates that the Rendition contains content which is considered essential to play.
	Forced *bool `json:"forced,omitempty"`
}

