package model

type EmailNotificationWithStreamConditionsRequest struct {
	// Notify when condition resolves after it was met
	Resolve *bool `json:"resolve,omitempty"`
	Emails []string `json:"emails,omitempty"`
	Muted *bool `json:"muted,omitempty"`
	Conditions *AbstractCondition `json:"conditions,omitempty"`
}

