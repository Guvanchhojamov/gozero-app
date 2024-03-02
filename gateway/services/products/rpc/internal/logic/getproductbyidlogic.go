package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByIdLogic {
	return &GetProductByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductByIdLogic) GetProductById(in *v1.GetProductByIdRequest) (*v1.GetProductByIdResponse, error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "CreateProductLogic.Create")
	defer span.End()
	product, err := l.svcCtx.App.Repository.Product.GetProductById(ctx, in)
	if err != nil {
		return nil, err
	}
	return product, nil
}
