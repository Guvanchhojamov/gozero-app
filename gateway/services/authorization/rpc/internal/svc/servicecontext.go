package svc

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/app"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	App    *app.App
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	newApp, err := app.NewApp(c)
	if err != nil {
		return nil, err
	}
	return &ServiceContext{
		Config: c,
		App:    newApp,
	}, nil
}
