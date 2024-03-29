package authorization

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/userauthservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInResp, err error) {
	input := &userauthservice.SignInRequest{
		Login:    req.Login,
		Password: req.Password,
	}
	response, err := l.svcCtx.Authorization.SignIn(l.ctx, input)
	if err != nil {
		//logx.Errorf("signInLogic err: %s", err)
		return nil, err
	}
	return &types.SignInResp{Token: response.Token}, err
}
