package app

import (
	"context"
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/models"
	"github.com/Guvanchhojamov/gozero-app/third_party/database"
	"github.com/go-errors/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zeromicro/go-zero/core/logx"
)

type Authorization struct {
	authJwtSignKey string
	db             *pgxpool.Pool
}

func NewAuthorization(authJwtSignKey string, dbParams config.Postgres) (*Authorization, error) {
	dbPool, err := database.NewPostgres(context.Background(), database.Params(dbParams))
	if err != nil {
		logx.Error(err)
		return nil, errors.New(err)
	}
	return &Authorization{
		authJwtSignKey: authJwtSignKey,
		db:             dbPool,
	}, nil
}

func (a *Authorization) ParseToken(accessToken string) (userID uint32, err error) {
	var claims models.CustomTokenClaims
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.authJwtSignKey), nil
	})
	if err != nil {
		return 0, err
	}
	switch {
	case !token.Valid:
		return 0, fmt.Errorf("errParseToken.invalidTokenMsg")
	case err != nil:
		return 0, fmt.Errorf("errParseToken.Unknown")
	}
	return claims.UserId, err
}

func (a *Authorization) CheckIsAdmin(ctx context.Context, userId uint32) (isAdmin bool, err error) {
	var roleId uint32
	var roleName string
	query := fmt.Sprintf(`SELECT 
    								roles.id, roles.name 
									FROM %[1]s
									    JOIN %[2]s ON roles.id = users.role_id  WHERE users.id = $1`, rolesTable, usersTable)

	err = a.db.QueryRow(ctx, query, userId).Scan(&roleId, &roleName)
	if err != nil {
		return false, errors.New(err)
	}
	if roleId != adminRoleId || roleName != adminRoleName {
		return false, ErrAccessDenied
	}
	return true, nil
}
func (a *Authorization) CheckRolePermission(ctx context.Context, userId uint32) (hasPermission bool, err error) {
	var roleId uint32
	var roleName string
	query := fmt.Sprintf(`SELECT 
    								roles.id, roles.name 
									FROM %[1]s
									    JOIN %[2]s ON roles.id = users.role_id  WHERE users.id = $1`, rolesTable, usersTable)

	err = a.db.QueryRow(ctx, query, userId).Scan(&roleId, &roleName)
	if err != nil {
		return false, errors.New(err)
	}
	if roleName == userRoleName {
		return false, ErrAccessDenied
	}
	return true, nil
}
