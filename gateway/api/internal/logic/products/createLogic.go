package products

import (
	"context"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateProductReq) (resp *types.CreateProductResp, err error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "CreateProductLogic.Create")
	defer span.End()
	input := &v1.CreateProductRequest{
		Name:  req.Name,
		Price: req.Price,
	}
	product, err := l.svcCtx.Product.CreateProduct(ctx, input)
	if err != nil {
		return nil, err
	}
	return &types.CreateProductResp{ProductId: product.ProductId}, nil
}
