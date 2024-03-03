package orders

import (
	"context"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"

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
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "GetOrders.GetOrders")
	defer span.End()
	resp, err = l.svcCtx.Order.GetOrders(ctx, &v1.GetOrdersRequest{})
	if err != nil {
		logx.ErrorStack(err)
		return nil, err
	}
	return resp, nil
}
