// Code generated by goctl. DO NOT EDIT.
// Source: auth.proto

package userauthservice

import (
	"context"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SignInRequest  = v1.SignInRequest
	SignInResponse = v1.SignInResponse
	SignUpRequest  = v1.SignUpRequest
	SignUpResponse = v1.SignUpResponse

	UserAuthService interface {
		SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
		SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	}

	defaultUserAuthService struct {
		cli zrpc.Client
	}
)

func NewUserAuthService(cli zrpc.Client) UserAuthService {
	return &defaultUserAuthService{
		cli: cli,
	}
}

func (m *defaultUserAuthService) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	client := v1.NewUserAuthServiceClient(m.cli.Conn())
	return client.SignUp(ctx, in, opts...)
}

func (m *defaultUserAuthService) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	client := v1.NewUserAuthServiceClient(m.cli.Conn())
	return client.SignIn(ctx, in, opts...)
}
