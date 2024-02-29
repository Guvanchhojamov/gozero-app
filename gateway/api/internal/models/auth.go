package models

import "github.com/golang-jwt/jwt/v4"

type ClientTokenClaims struct {
	UserId uint32 `json:"userId"`
	jwt.RegisteredClaims
}
