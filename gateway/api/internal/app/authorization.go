package app

import (
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

type Authorization struct {
	c config.Config
}

func NewAuthorization(c config.Config) *Authorization {
	return &Authorization{
		c: c,
	}
}

func (a *Authorization) ParseToken(accessToken string) (userID uint32, err error) {
	var claims models.CustomTokenClaims
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.c.App.JWTSecretKey), nil
	})
	if err != nil {
		return 0, err
	}
	switch {
	case !token.Valid:
		return 0, fmt.Errorf("errParseToken.invalidTokenMsg")
	case err != nil:
		return 0, fmt.Errorf("errParseToken.Unknown")
	}
	return claims.UserId, err
}
