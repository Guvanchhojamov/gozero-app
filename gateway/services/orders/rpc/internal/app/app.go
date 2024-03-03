package app

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/repository"
	"github.com/Guvanchhojamov/gozero-app/third_party/database"
	"github.com/go-errors/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type App struct {
	Repository *repository.Repository
}

func NewApp(cnf config.Config) (*App, error) {
	dbPoll, err := database.NewPostgres(context.Background(), database.Params(cnf.App.Postgres))
	if err != nil {
		logx.Error(err)
		return nil, errors.New(err)
	}
	return &App{Repository: repository.NewRepository(dbPoll)}, nil
}
