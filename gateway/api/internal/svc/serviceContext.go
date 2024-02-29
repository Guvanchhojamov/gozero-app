package svc

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/app"
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
	App                        *app.Domain
}

func NewServiceContext(c config.Config, appDomain *app.Domain) *ServiceContext {
	return &ServiceContext{
		Config:                     c,
		RolePermissionMiddleware:   middleware.NewRolePermissionMiddleware(appDomain.Authorization).Handle,
		HeaderValidationMiddleware: middleware.NewHeaderValidationMiddleware(appDomain.Authorization).Handle,
		Authorization:              userauthservice.NewUserAuthService(zrpc.MustNewClient(c.Services.Authorization)),
		App:                        app.NewDomain(c),
	}
}
