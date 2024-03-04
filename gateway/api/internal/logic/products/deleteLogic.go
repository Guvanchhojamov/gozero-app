package products

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteProductReq) (resp *types.DeleteProductResp, err error) {
	input := &v1.DeleteProductRequest{ProductId: req.ProductId}
	response, err := l.svcCtx.Product.DeleteProduct(l.ctx, input)
	if err != nil {
		return nil, err
	}
	resp = &types.DeleteProductResp{
		StatusCode: response.StatusCode,
		Message:    response.Message,
	}
	return resp, nil
}
