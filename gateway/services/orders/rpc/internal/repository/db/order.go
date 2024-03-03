package db

import (
	"context"
	"fmt"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{db: db}
}
func (o *OrderRepo) GetOrders(ctx context.Context, input *v1.GetOrdersRequest) (*v1.GetOrdersResponse, error) {
	fmt.Println("in order db")
	return nil, nil
}
