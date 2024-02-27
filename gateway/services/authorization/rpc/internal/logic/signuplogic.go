package logic

import (
	"context"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpLogic) SignUp(in *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.SignUpResponse{}, nil
}
