package orders

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

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
	input := &v1.UpdateOrderRequest{
		Id:        req.Id,
		UserId:    &req.UserId,
		ProductId: &req.ProductId,
		Price:     req.Price,
	}
	order, err := l.svcCtx.Order.UpdateOrder(l.ctx, input)
	if err != nil {
		return nil, err
	}
	return &types.UpdateOrderResp{OrderId: order.OrderId}, nil
}
