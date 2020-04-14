package model
import (
	"time"
)

type Webhook struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Webhook url (required)
	Url string `json:"url,omitempty"`
	// HTTP method used for the webhook
	Method WebhookHttpMethod `json:"method,omitempty"`
	// Whether to skip SSL certification verification or not
	InsecureSsl *bool `json:"insecureSsl,omitempty"`
	// Encryption used for the webhook
	Encryption *WebhookEncryption `json:"encryption,omitempty"`
	// Signature used for the webhook
	Signature *WebhookSignature `json:"signature,omitempty"`
}

