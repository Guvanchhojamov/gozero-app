// Code generated by goctl. DO NOT EDIT.
// Source: orders.proto

package server

import (
	"context"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/logic"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"
)

type OrderServiceServer struct {
	svcCtx *svc.ServiceContext
	v1.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(svcCtx *svc.ServiceContext) *OrderServiceServer {
	return &OrderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServiceServer) GetOrders(ctx context.Context, in *v1.GetOrdersRequest) (*v1.GetOrdersResponse, error) {
	l := logic.NewGetOrdersLogic(ctx, s.svcCtx)
	return l.GetOrders(in)
}

func (s *OrderServiceServer) GetOrderById(ctx context.Context, in *v1.GetOrderByIdRequest) (*v1.GetOrderByIdResponse, error) {
	l := logic.NewGetOrderByIdLogic(ctx, s.svcCtx)
	return l.GetOrderById(in)
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, in *v1.CreateOrderRequest) (*v1.CreateOrderResponse, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *OrderServiceServer) UpdateOrder(ctx context.Context, in *v1.UpdateOrderRequest) (*v1.UpdateOrderResponse, error) {
	l := logic.NewUpdateOrderLogic(ctx, s.svcCtx)
	return l.UpdateOrder(in)
}

func (s *OrderServiceServer) DeleteOrder(ctx context.Context, in *v1.DeleteOrderRequest) (*v1.DeleteOrderResponse, error) {
	l := logic.NewDeleteOrderLogic(ctx, s.svcCtx)
	return l.DeleteOrder(in)
}
