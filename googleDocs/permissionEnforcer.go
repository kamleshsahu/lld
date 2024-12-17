package main

import (
	"errors"
	"fmt"
)

// Permission constants
type Permission string

const (
	CanCreate Permission = "create"
	CanUpdate Permission = "update"
	CanDelete Permission = "delete"
	CanView   Permission = "view"
)

// PermissionEnforcer checks RBAC
type PermissionEnforcer struct{}

func (pe PermissionEnforcer) Enforce(user *User, permission Permission) error {
	if allowed, exists := user.Role.Permissions[permission]; exists && allowed {
		return nil
	}
	return errors.New(fmt.Sprintf("permission denied: user %s cannot %s", user.Name, permission))
}
