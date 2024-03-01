package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &v1.GetProductByIdResponse{}, nil
}
