package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrdersLogic {
	return &GetOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrdersLogic) GetOrders(in *v1.GetOrdersRequest) (*v1.GetOrdersResponse, error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "CreateOrdersLogic.GeOrders")
	defer span.End()
	ordersResp, err := l.svcCtx.App.Repository.Order.GetOrders(ctx, in)
	if err != nil {
		return nil, err
	}
	return ordersResp, nil
}
