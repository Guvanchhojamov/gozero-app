package orders

import (
	"context"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"

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

func (l *GetOrderByIdLogic) GetOrderById(req *types.GetOrderByIdReq) (resp *types.GetOrderByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
