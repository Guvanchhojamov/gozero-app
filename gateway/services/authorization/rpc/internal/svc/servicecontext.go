package svc

import "github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
