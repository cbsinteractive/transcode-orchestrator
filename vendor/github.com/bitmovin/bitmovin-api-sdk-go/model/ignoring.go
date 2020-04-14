package model

type Ignoring struct {
	IgnoredBy IgnoredBy `json:"ignoredBy,omitempty"`
	// Describes why ignoredBy has been set to its current value.
	IgnoredByDescription string `json:"ignoredByDescription,omitempty"`
}

