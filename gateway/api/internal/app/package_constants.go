package app

import "fmt"

const (
	adminRoleName   = "admin"
	adminRoleId     = 1
	rolesTable      = "roles"
	usersTable      = "users"
	sellerRoleName  = "seller"
	sellerRoleId    = 2
	deliverRoleName = "deliver"
	deliverRoleId   = 3
	userRoleName    = "user"
)

var (
	ErrAccessDenied = fmt.Errorf("%s", "USER_ACCESS_DENIED")
)
