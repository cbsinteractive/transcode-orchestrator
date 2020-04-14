package model

type WebhookEncryption struct {
	// The encryption algorithm used for the webhook (required)
	Type EncryptionType `json:"type,omitempty"`
	// The key of the encryption (required)
	Key string `json:"key,omitempty"`
}

