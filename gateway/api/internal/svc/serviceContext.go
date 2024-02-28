package svc

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                     config.Config
	HeaderValidationMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                     c,
		HeaderValidationMiddleware: middleware.NewHeaderValidationMiddleware().Handle,
	}
}
