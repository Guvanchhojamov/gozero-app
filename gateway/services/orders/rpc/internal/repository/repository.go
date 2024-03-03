package repository

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/repository/db"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IOrder interface {
	GetOrders(ctx context.Context, input *v1.GetOrdersRequest) (*v1.GetOrdersResponse, error)
}

type Repository struct {
	Order IOrder
}

func NewRepository(dbPoll *pgxpool.Pool) *Repository {
	return &Repository{Order: db.NewOrderRepo(dbPoll)}
}
