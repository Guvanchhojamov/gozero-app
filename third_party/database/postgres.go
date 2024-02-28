package database

import (
	"context"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/url"
)

// Params - database connection params
type Params struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	SslMode  string
}

// NewPostgres - establishing database connection
func NewPostgres(ctx context.Context, cfg Params) (*pgxpool.Pool, error) {
	dbURL := "postgres://" + cfg.Username + ":" + url.QueryEscape(cfg.Password) + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DbName + "?sslmode=" + cfg.SslMode
	//fmt.Println(dbURL)
	dbPool, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, errors.New(err)
	}
	err = dbPool.Ping(ctx)
	if err != nil {
		return nil, errors.New(err)
	}
	return dbPool, nil
}
