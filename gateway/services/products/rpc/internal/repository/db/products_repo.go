package db

import (
	"context"
	"fmt"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) CreateProduct(ctx context.Context, input *v1.CreateProductRequest) (*v1.CreateProductResponse, error) {

	var id uint32
	query := fmt.Sprintf("select id from users where users.id=1")

	err := p.db.QueryRow(ctx, query).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	return &v1.CreateProductResponse{}, nil
}
