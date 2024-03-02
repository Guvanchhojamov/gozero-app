package products

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdLogic) GetById(req *types.GetProductByIdReq) (resp *v1.GetProductByIdResponse, err error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "CreateProductLogic.Create")
	defer span.End()
	input := &v1.GetProductByIdRequest{Id: req.ProductId}
	resp, err = l.svcCtx.Product.GetProductById(ctx, input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
