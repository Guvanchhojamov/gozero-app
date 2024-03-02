package products

import (
	"context"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductsLogic {
	return &GetProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductsLogic) GetProducts(req *types.GetProductsReq) (resp *v1.GetProductsResponse, err error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "GetProducts.GetProducts")
	defer span.End()

	products, err := l.svcCtx.Product.GetProducts(ctx, &v1.GetProductsRequest{})
	if err != nil {
		return nil, err
	}
	return &v1.GetProductsResponse{Products: products.Products}, nil
}
