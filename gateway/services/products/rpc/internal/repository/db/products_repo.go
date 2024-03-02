package db

import (
	"context"
	"fmt"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) CreateProduct(ctx context.Context, input *v1.CreateProductRequest) (*v1.CreateProductResponse, error) {
	var prodId uint32
	query := fmt.Sprintf(`INSERT INTO %[1]s (name, price)
										VALUES ($1, $2) RETURNING id`, productsTable)
	err := p.db.QueryRow(ctx, query, input.Name, input.Price).Scan(&prodId)
	if err != nil {
		return &v1.CreateProductResponse{}, errors.New(err)
	}
	return &v1.CreateProductResponse{ProductId: prodId}, nil
}

func (p *ProductRepo) GetProducts(ctx context.Context, input *v1.GetProductsRequest) (*v1.GetProductsResponse, error) {
	var products []*v1.Product
	query := fmt.Sprintf(`SELECT id,name,price  
									FROM %s`, productsTable)
	rows, err := p.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, errors.New(err)
	}
	for rows.Next() {
		var product v1.Product
		if err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return &v1.GetProductsResponse{Products: products}, nil
}
