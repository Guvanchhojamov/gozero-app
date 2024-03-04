package repository

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/third_party/database"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(cnf config.Config) (*AuthRepository, error) {
	dbParams := database.Params(cnf.App.DateBase.Postgres)
	dbPool, err := database.NewPostgres(context.Background(), dbParams)
	if err != nil {
		logx.Error(err)
		return nil, errors.New(err)
	}
	return &AuthRepository{
		db: dbPool,
	}, nil
}
