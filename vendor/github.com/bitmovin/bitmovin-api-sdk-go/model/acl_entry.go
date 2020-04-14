package model

type AclEntry struct {
	Scope string `json:"scope,omitempty"`
	Permission AclPermission `json:"permission,omitempty"`
}

