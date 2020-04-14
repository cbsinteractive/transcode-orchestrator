package model

type WebhookSignature struct {
	// The signature type used for the webhook.  Selects one of the supported signatures. The signature is attached to the list of headers with the key `Bitmovin-Signature`. In case of the `HMAC` type the SHA512 hashing algorithm is used to generate an authentication code from the webhook body. (required)
	Type SignatureType `json:"type,omitempty"`
	// The key of the signature (required)
	Key string `json:"key,omitempty"`
}

