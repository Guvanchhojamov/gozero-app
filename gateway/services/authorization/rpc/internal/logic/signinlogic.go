package logic

import (
	"context"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *v1.SignInRequest) (*v1.SignUpResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.SignUpResponse{}, nil
}
