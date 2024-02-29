package app

import (
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

type Authorization struct {
	authJwtSignKey string
}

func NewAuthorization(authJwtSignKey string) *Authorization {
	return &Authorization{authJwtSignKey: authJwtSignKey}
}

func (a *Authorization) ParseToken(accessToken string) (userID uint32, err error) {
	var claims models.ClientTokenClaims
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.authJwtSignKey), nil
	})
	switch {
	case !token.Valid:
		return 0, fmt.Errorf("errParseToken.invalidTokenMsg")
	case err != nil:
		return 0, fmt.Errorf("errParseToken.Unknown")
	}
	return claims.UserId, err
}
