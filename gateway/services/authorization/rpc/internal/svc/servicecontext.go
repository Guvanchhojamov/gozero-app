package svc

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/domain"
)

type ServiceContext struct {
	Config config.Config
	App    *domain.App
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	newApp, err := domain.NewApp(c.App)
	if err != nil {
		return nil, err
	}
	return &ServiceContext{
		Config: c,
		App:    newApp,
	}, nil
}
