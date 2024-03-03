package svc

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/app"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/middleware"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/userauthservice"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/orderservice"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/productservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                     config.Config
	RolePermissionMiddleware   rest.Middleware
	HeaderValidationMiddleware rest.Middleware
	IsAdminValidateMiddleware  rest.Middleware
	App                        *app.Domain
	Authorization              userauthservice.UserAuthService
	Product                    productservice.ProductService
	Order                      orderservice.OrderService
}

func NewServiceContext(c config.Config, appDomain *app.Domain) *ServiceContext {
	return &ServiceContext{
		Config:                     c,
		RolePermissionMiddleware:   middleware.NewRolePermissionMiddleware(appDomain.Authorization).Handle,
		HeaderValidationMiddleware: middleware.NewHeaderValidationMiddleware(appDomain.Authorization).Handle,
		IsAdminValidateMiddleware:  middleware.NewIsAdminValidateMiddleware(appDomain.Authorization).Handle,
		App:                        app.NewDomain(c),
		Authorization:              userauthservice.NewUserAuthService(zrpc.MustNewClient(c.Services.Authorization)),
		Product:                    productservice.NewProductService(zrpc.MustNewClient(c.Services.Product)),
		Order:                      orderservice.NewOrderService(zrpc.MustNewClient(c.Services.Order)),
	}
}
