package repository

import (
	"context"
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/userauthservice"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, userInput *userauthservice.SignUpRequest) (userId uint32, err error) {
	query := fmt.Sprintf(`INSERT INTO users (login,role_id, password) VALUES ($1, $2, $3, ...)`)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return 0, err
	}
	fmt.Println(rows)
	return 0, nil
}
