package logic

import (
	"context"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *v1.DeleteProductRequest) (*v1.DeleteProductResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.DeleteProductResponse{}, nil
}
