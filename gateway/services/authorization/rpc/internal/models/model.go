package models

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id uint32
}

type CustomTokenClaims struct {
	jwt.RegisteredClaims
	UserId uint32 `json:"userId"`
}
