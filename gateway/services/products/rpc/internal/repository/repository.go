package repository

import (
	"context"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/internal/repository/db"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IProduct interface {
	CreateProduct(ctx context.Context, input *v1.CreateProductRequest) (*v1.CreateProductResponse, error)
}

type Repository struct {
	Product IProduct
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{Product: db.NewProductRepo(dbPool)}
}