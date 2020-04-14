package model

type CencWidevine struct {
	// Base64 encoded pssh payload (required)
	Pssh string `json:"pssh,omitempty"`
}

