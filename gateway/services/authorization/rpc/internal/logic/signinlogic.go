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

func (l *SignInLogic) SignIn(in *v1.SignInRequest) (*v1.SignInResponse, error) {
	token, err := l.svcCtx.App.Repository.GenerateToken(l.ctx, in.Login, in.Password)
	if err != nil {
		logx.Errorf("errGenerateToken: %s", err)
		return nil, err
	}

	return &v1.SignInResponse{Token: token}, nil
}
