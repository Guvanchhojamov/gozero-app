package middleware

import "fmt"

// Header validation middleware
const (
	authHeaderKey   = "Authorization"
	userIdCtxKey    = "userId"
	errUnauthorized = "error: Unauthorized"
	roleIdCtxKey    = "roleId"
)

var (
	ErrAccessDenied = fmt.Errorf("%s", "USER_ACCESS_DENIED")
)
