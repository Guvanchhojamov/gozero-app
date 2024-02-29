package app

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/domain"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type Domain struct {
	domain.Authorization
}

func NewDomain(cnf config.Config) *Domain {
	newAuth, err := NewAuthorization(cnf.App.JWTSecretKey, cnf.App.DateBase.Postgres)
	if err != nil {
		logx.ErrorStack(err)
		return nil
	}
	return &Domain{
		Authorization: newAuth,
	}
}
