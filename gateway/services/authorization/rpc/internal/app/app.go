package app

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/repository"
	"github.com/Guvanchhojamov/gozero-app/third_party/database"
	"github.com/zeromicro/go-zero/core/logx"
)

type App struct {
	Repository *repository.Repository
}

func NewApp(c config.Config) (*App, error) {
	dbPool, err := database.NewPostgres(context.Background(), database.Params(c.App.Postgres))
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	newRepo := repository.NewRepository(dbPool)
	return &App{Repository: newRepo}, nil
}
