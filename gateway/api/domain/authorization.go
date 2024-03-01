package domain

import "context"

type Authorization interface {
	ParseToken(accessToken string) (userId uint32, err error)
	CheckIsAdmin(ctx context.Context, userId uint32) (isAdmin bool, err error)
	CheckRolePermission(ctx context.Context, userId uint32) (hasPermission bool, err error)
}
