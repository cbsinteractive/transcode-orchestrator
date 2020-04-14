package model
// InvitationStatus : Status of the invitation
type InvitationStatus string

// List of InvitationStatus
const (
	InvitationStatus_PENDING InvitationStatus = "PENDING"
	InvitationStatus_ACCEPTED InvitationStatus = "ACCEPTED"
)