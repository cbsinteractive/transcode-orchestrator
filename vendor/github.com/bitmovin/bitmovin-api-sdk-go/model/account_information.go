package model
import (
	"time"
)

type AccountInformation struct {
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
	// Email address of the account. (required)
	Email string `json:"email,omitempty"`
	// ApiKeys associated with the account (required)
	ApiKeys []AccountApiKey `json:"apiKeys,omitempty"`
	// First name of the tenant.
	FirstName string `json:"firstName,omitempty"`
	// Last name of the tenant.
	LastName string `json:"lastName,omitempty"`
	// Phone number of the tenant.
	Phone string `json:"phone,omitempty"`
	// Company name of the tenant.
	Company string `json:"company,omitempty"`
	IsVerified *bool `json:"isVerified,omitempty"`
}

