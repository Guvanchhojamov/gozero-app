package models

import "github.com/golang-jwt/jwt/v4"

type CustomTokenClaims struct {
	UserId uint32 `json:"userId"`
	jwt.RegisteredClaims
}
