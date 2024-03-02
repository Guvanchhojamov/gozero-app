package logic

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *v1.CreateProductRequest) (*v1.CreateProductResponse, error) {
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "CreateProductLogic.Create")
	defer span.End()
	input := &v1.CreateProductRequest{
		Name:  in.Name,
		Price: in.Price,
	}
	newProduct, err := l.svcCtx.App.Repository.Product.CreateProduct(ctx, input)
	if err != nil {
		return nil, err
	}
	return &v1.CreateProductResponse{ProductId: newProduct.ProductId}, nil
}
