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
	App                        *app.App
	Authorization              userauthservice.UserAuthService
	Product                    productservice.ProductService
	Order                      orderservice.OrderService
}

func NewServiceContext(c config.Config, newApp *app.App) *ServiceContext {
	return &ServiceContext{
		Config:                     c,
		RolePermissionMiddleware:   middleware.NewRolePermissionMiddleware(newApp.AuthRepo).Handle,
		HeaderValidationMiddleware: middleware.NewHeaderValidationMiddleware(newApp.Authorization).Handle,
		IsAdminValidateMiddleware:  middleware.NewIsAdminValidateMiddleware(newApp.AuthRepo).Handle,
		App:                        newApp,
		Authorization:              userauthservice.NewUserAuthService(zrpc.MustNewClient(c.Services.Authorization)),
		Product:                    productservice.NewProductService(zrpc.MustNewClient(c.Services.Product)),
		Order:                      orderservice.NewOrderService(zrpc.MustNewClient(c.Services.Order)),
	}
}
