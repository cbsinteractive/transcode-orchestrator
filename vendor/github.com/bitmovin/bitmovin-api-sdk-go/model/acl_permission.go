package model
type AclPermission string

// List of AclPermission
const (
	AclPermission_PUBLIC_READ AclPermission = "PUBLIC_READ"
	AclPermission_PRIVATE AclPermission = "PRIVATE"
)