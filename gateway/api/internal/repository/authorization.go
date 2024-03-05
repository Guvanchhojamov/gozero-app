package repository

import (
	"context"
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/third_party/database"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(cnf config.Config) (*AuthRepository, error) {
	dbParams := database.Params(cnf.App.DateBase.Postgres)
	dbPool, err := database.NewPostgres(context.Background(), dbParams)
	if err != nil {
		logx.Error(err)
		return nil, errors.New(err)
	}
	return &AuthRepository{
		db: dbPool,
	}, nil
}
func (a *AuthRepository) CheckIsAdmin(ctx context.Context, userId uint32) (isAdmin bool, err error) {
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
func (a *AuthRepository) CheckRolePermission(ctx context.Context, userId uint32) (hasPermission bool, err error) {
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
