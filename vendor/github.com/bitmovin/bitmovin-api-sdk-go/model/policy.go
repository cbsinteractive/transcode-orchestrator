package model
// Policy : The policy for the permissions.
type Policy string

// List of Policy
const (
	Policy_ALLOW Policy = "ALLOW"
	Policy_DENY Policy = "DENY"
)