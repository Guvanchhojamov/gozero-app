package app

import "fmt"

const (
	adminRoleName = "admin"
	adminRoleId   = 1
	rolesTable    = "roles"
	usersTable    = "users"
)

var (
	ErrAccessDenied = fmt.Errorf("%s", "USER_ACCESS_DENIED")
)
