package logic

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *v1.CreateOrderRequest) (*v1.CreateOrderResponse, error) {
	order, err := l.svcCtx.App.Repository.Order.CreateOrder(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return order, nil
}
