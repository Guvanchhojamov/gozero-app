package app

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/domain"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
)

type Domain struct {
	domain.Authorization
}

func NewDomain(cnf config.Config) *Domain {
	return &Domain{
		Authorization: NewAuthorization(cnf.App.JWTSecretKey),
	}
}
