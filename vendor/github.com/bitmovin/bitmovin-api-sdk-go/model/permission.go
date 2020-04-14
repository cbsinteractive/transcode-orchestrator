package model
// Permission : Permission to assign.
type Permission string

// List of Permission
const (
	Permission_GET Permission = "GET"
	Permission_DELETE Permission = "DELETE"
	Permission_POST Permission = "POST"
	Permission_PUT Permission = "PUT"
	Permission_PATCH Permission = "PATCH"
)