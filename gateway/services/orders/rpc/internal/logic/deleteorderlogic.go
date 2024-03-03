package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOrderLogic) DeleteOrder(in *v1.DeleteOrderRequest) (*v1.DeleteOrderResponse, error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "deleteOrderLogic.Delete")
	defer span.End()
	order, err := l.svcCtx.App.Repository.Order.DeleteOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	return order, nil
}
