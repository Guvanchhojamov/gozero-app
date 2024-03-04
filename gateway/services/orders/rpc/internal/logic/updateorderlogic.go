package logic

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *v1.UpdateOrderRequest) (*v1.UpdateOrderResponse, error) {
	result, err := l.svcCtx.App.Repository.Order.UpdateOrder(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}
