package model

type CencFairPlay struct {
	// Initialization vector as hexadecimal string
	Iv string `json:"iv,omitempty"`
	// URL of the licensing server
	Uri string `json:"uri,omitempty"`
}

