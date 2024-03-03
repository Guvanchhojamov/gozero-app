package orders

import (
	"context"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderLogic) UpdateOrder(req *types.UpdateOrderReq) (resp *types.UpdateOrderResp, err error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "Order.Update")
	defer span.End()
	input := &v1.UpdateOrderRequest{
		Id:        req.Id,
		UserId:    &req.UserId,
		ProductId: &req.ProductId,
		Price:     req.Price,
	}
	order, err := l.svcCtx.Order.UpdateOrder(ctx, input)
	if err != nil {
		return nil, err
	}
	return &types.UpdateOrderResp{OrderId: order.OrderId}, nil
}
