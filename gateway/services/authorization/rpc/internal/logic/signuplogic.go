package logic

import (
	"context"
	"fmt"

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
	fmt.Println("in rpc logic")
	userId, err := l.svcCtx.App.Repository.CreateUser(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.SignUpResponse{UserId: userId}, nil
}
