package model

type CencPlayReady struct {
	// Url of the license server. Either the laUrl or the pssh needs to be provided.
	LaUrl string `json:"laUrl,omitempty"`
	// Base64 encoded pssh payload.
	Pssh string `json:"pssh,omitempty"`
}

