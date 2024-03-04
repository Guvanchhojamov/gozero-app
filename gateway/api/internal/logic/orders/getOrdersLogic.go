package orders

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrdersLogic {
	return &GetOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrdersLogic) GetOrders(req *types.GetOrdersReq) (resp *v1.GetOrdersResponse, err error) {
	resp, err = l.svcCtx.Order.GetOrders(l.ctx, &v1.GetOrdersRequest{})
	if err != nil {
		logx.ErrorStack(err)
		return nil, err
	}
	return resp, nil
}
