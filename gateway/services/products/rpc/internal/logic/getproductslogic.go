package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductsLogic {
	return &GetProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductsLogic) GetProducts(in *v1.GetProductsRequest) (*v1.GetProductsResponse, error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "CreateProductLogic.Create")
	defer span.End()
	products, err := l.svcCtx.App.Repository.Product.GetProducts(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.GetProductsResponse{Products: products.Products}, nil
}
