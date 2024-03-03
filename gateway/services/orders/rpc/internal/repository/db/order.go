package db

import (
	"context"
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/common/response"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"net/http"
	"strings"
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
func (o *OrderRepo) GetOrderById(ctx context.Context, input *v1.GetOrderByIdRequest) (*v1.GetOrderByIdResponse, error) {
	var order v1.OrderById
	order.User = &v1.OrderUser{}
	order.Product = &v1.OrderProduct{}
	query := fmt.Sprintf(`SELECT
    								o.id ,
    								o.price,
    								u.id as userId,
    								u.login as username,
    								p.id as productId,
    								p.name
            							FROM %[1]s  o
            							JOIN %[2]s u ON u.id = o.user_id
            							JOIN %[3]s p on p.id = o.product_id WHERE o.id=$1`, ordersTable, usersTable, productsTable)

	if err := o.db.QueryRow(ctx, query, input.Id).Scan(
		&order.Id,
		&order.Price,
		&order.User.UserId,
		&order.User.Username,
		&order.Product.ProductId,
		&order.Product.Name,
	); err != nil {
		return nil, response.NewRpcErrResponse(codes.Internal, errors.New(err))
	}
	return &v1.GetOrderByIdResponse{Order: &order}, nil
}

func (o *OrderRepo) CreateOrder(ctx context.Context, input *v1.CreateOrderRequest) (*v1.CreateOrderResponse, error) {
	var ordId uint32
	query := fmt.Sprintf(`INSERT INTO %[1]s (user_id, product_id,price)
										VALUES ($1, $2,$3) RETURNING id`, ordersTable)
	err := o.db.QueryRow(ctx, query, input.UserId, input.ProductId, input.Price).Scan(&ordId)
	if err != nil {
		return &v1.CreateOrderResponse{}, response.NewRpcErrResponse(codes.Internal, errors.New(err))
	}
	return &v1.CreateOrderResponse{OrderId: ordId}, nil
}

func (o *OrderRepo) DeleteOrder(ctx context.Context, input *v1.DeleteOrderRequest) (*v1.DeleteOrderResponse, error) {
	var resp *v1.DeleteOrderResponse
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, ordersTable)
	exec, err := o.db.Exec(ctx, query, input.OrderId)
	if err != nil {
		return nil, response.NewRpcErrResponse(codes.Internal, errors.New(err))
	}
	if rowsAffected := exec.RowsAffected(); rowsAffected > 0 {
		resp = &v1.DeleteOrderResponse{
			StatusCode: http.StatusOK,
			Message:    "success",
		}
		return resp, nil
	}
	resp = &v1.DeleteOrderResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "error",
	}
	return resp, nil
}
func (o *OrderRepo) UpdateOrder(ctx context.Context, input *v1.UpdateOrderRequest) (*v1.UpdateOrderResponse, error) {
	var setVals = make([]string, 0)
	var args = make([]interface{}, 0)
	var argId = 1
	fmt.Println(input)
	if *input.UserId != 0 {
		setVals = append(setVals, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *input.UserId)
		argId++
	}

	if *input.ProductId != 0 {
		setVals = append(setVals, fmt.Sprintf("product_id=$%d", argId))
		args = append(args, *input.ProductId)
		argId++
	}
	if input.Price != 0 {
		setVals = append(setVals, fmt.Sprintf("price=$%d", argId))
		args = append(args, input.Price)
		argId++
	}

	setString := strings.Join(setVals, ", ")
	args = append(args, input.Id)
	updateQuery := fmt.Sprintf(`UPDATE %[1]s 
												SET %[2]s
													WHERE id=$%[3]d`, ordersTable, setString, argId)
	fmt.Println(updateQuery)
	fmt.Println(args)
	exec, err := o.db.Exec(ctx, updateQuery, args...)
	if err != nil {
		return nil, response.NewRpcErrResponse(codes.Internal, errors.New(err))
	}
	if rowsAffected := exec.RowsAffected(); rowsAffected <= 0 {
		return nil, response.NewRpcErrResponse(codes.Internal, errors.New("order not found"))
	}
	return &v1.UpdateOrderResponse{OrderId: input.Id}, nil
}
