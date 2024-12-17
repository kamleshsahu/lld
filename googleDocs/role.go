package main

// Role defines a user's role and associated permissions
type Role struct {
	Name        string
	Permissions map[Permission]bool
}

var AdminRole = Role{Name: "admin", Permissions: map[Permission]bool{CanCreate: true, CanUpdate: true, CanDelete: true, CanView: true}}
var EditorRole = Role{Name: "editor", Permissions: map[Permission]bool{CanCreate: true, CanUpdate: true, CanDelete: false, CanView: true}}
var ViewerRole = Role{Name: "viewer", Permissions: map[Permission]bool{CanCreate: false, CanUpdate: false, CanDelete: false, CanView: true}}
