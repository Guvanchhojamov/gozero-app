package orders

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *v1.CreateOrderResponse, err error) {
	input := &v1.CreateOrderRequest{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Price:     req.Price,
	}
	resp, err = l.svcCtx.Order.CreateOrder(l.ctx, input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
