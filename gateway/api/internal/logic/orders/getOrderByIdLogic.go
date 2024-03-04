package orders

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderByIdLogic) GetOrderById(req *types.GetOrderByIdReq) (resp *v1.GetOrderByIdResponse, err error) {
	resp, err = l.svcCtx.Order.GetOrderById(l.ctx, &v1.GetOrderByIdRequest{Id: req.OrderId})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
