package svc

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/middleware"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/userauthservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                     config.Config
	RolePermissionMiddleware   rest.Middleware
	HeaderValidationMiddleware rest.Middleware
	Authorization              userauthservice.UserAuthService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                     c,
		RolePermissionMiddleware:   middleware.NewRolePermissionMiddleware().Handle,
		HeaderValidationMiddleware: middleware.NewHeaderValidationMiddleware().Handle,
		Authorization:              userauthservice.NewUserAuthService(zrpc.MustNewClient(c.Services.Authorization)),
	}
}
