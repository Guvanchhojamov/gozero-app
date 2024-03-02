package products

import (
	"context"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"

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
	ctx, span := trace.TracerFromContext(l.ctx).Start(l.ctx, "UpdateProducts.Update")
	defer span.End()
	input := &v1.UpdateProductRequest{
		Id:    req.Id,
		Name:  &req.Name,
		Price: &req.Price,
	}
	resp, err = l.svcCtx.Product.UpdateProduct(ctx, input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
