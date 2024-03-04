package products

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateProductReq) (resp *v1.UpdateProductResponse, err error) {
	input := &v1.UpdateProductRequest{
		Id:    req.Id,
		Name:  &req.Name,
		Price: &req.Price,
	}
	resp, err = l.svcCtx.Product.UpdateProduct(l.ctx, input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
