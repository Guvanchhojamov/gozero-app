package db

import (
	"context"
	"fmt"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{db: db}
}
func (o *OrderRepo) GetOrders(ctx context.Context, input *v1.GetOrdersRequest) (*v1.GetOrdersResponse, error) {
	var orders []*v1.Order
	query := fmt.Sprintf(`SELECT id,user_id,product_id,price
									FROM %s`, ordersTable)
	rows, err := o.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, errors.New(err)
	}
	for rows.Next() {
		var order v1.Order
		if err = rows.Scan(
			&order.Id,
			&order.ProductId,
			&order.UserId,
			&order.Price,
		); err != nil {
			return nil, errors.New(err)
		}
		orders = append(orders, &order)
	}
	return &v1.GetOrdersResponse{Orders: orders}, nil
}
func (o *OrderRepo) GetOrderById(ctx context.Context, input *v1.GetOrdersRequest) (*v1.GetOrdersResponse, error) {
	var orders []*v1.Order
	query := fmt.Sprintf(`SELECT id,user_id,product_id,price
									FROM %s`, ordersTable)
	rows, err := o.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, errors.New(err)
	}
	for rows.Next() {
		var order v1.Order
		if err = rows.Scan(
			&order.Id,
			&order.ProductId,
			&order.UserId,
			&order.Price,
		); err != nil {
			return nil, errors.New(err)
		}
		orders = append(orders, &order)
	}
	return &v1.GetOrdersResponse{Orders: orders}, nil
}
