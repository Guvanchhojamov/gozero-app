package middleware

import (
	"fmt"
	"net/http"
)

type HeaderValidationMiddleware struct {
}

func NewHeaderValidationMiddleware() *HeaderValidationMiddleware {
	return &HeaderValidationMiddleware{}
}

func (m *HeaderValidationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		fmt.Println("from middeware")
		// Passthrough to next handler if need
		next(w, r)
	}
}
