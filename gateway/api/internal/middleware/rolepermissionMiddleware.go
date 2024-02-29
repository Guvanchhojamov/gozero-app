package middleware

import (
	"fmt"
	"net/http"
)

type RolePermissionMiddleware struct {
}

func NewRolePermissionMiddleware() *RolePermissionMiddleware {
	return &RolePermissionMiddleware{}
}

func (m *RolePermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		fmt.Println("role middleware")
		// Passthrough to next handler if need
		next(w, r)
	}
}
