package authorization

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/userauthservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	signUpReq := &userauthservice.SignUpRequest{
		Login:    req.Login,
		Password: req.Password,
		RoleId:   req.RoleId,
	}
	signUpResp, err := l.svcCtx.Authorization.SignUp(l.ctx, signUpReq)
	if err != nil {
		return nil, err
	}
	return &types.SignUpResp{UserId: signUpResp.UserId}, err
}
