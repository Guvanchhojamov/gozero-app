package app

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/repository"
	"github.com/go-errors/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type App struct {
	AuthRepo      repository.AuthRepository
	Authorization Authorization
}

func NewApp(c config.Config) (*App, error) {
	newAuthRepo, err := repository.NewAuthRepository(c)
	if err != nil {
		logx.Error(err)
		return nil, errors.New(err)
	}
	return &App{
		AuthRepo:      *newAuthRepo,
		Authorization: *NewAuthorization(c),
	}, nil
}
