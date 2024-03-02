package db

import (
	"context"
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/common/response"
	v1 "github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"net/http"
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
		return &v1.CreateProductResponse{}, response.NewRpcErrResponse(codes.Internal, errors.New(err))
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
		return nil, response.NewRpcErrResponse(codes.Internal, errors.New(err))
	}
	for rows.Next() {
		var product v1.Product
		if err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price); err != nil {
			return nil, errors.New(err)
		}
		products = append(products, &product)
	}
	return &v1.GetProductsResponse{Products: products}, nil
}
func (p *ProductRepo) GetProductById(ctx context.Context, input *v1.GetProductByIdRequest) (*v1.GetProductByIdResponse, error) {
	var product v1.Product
	query := fmt.Sprintf(`SELECT id,name,price
									FROM %s WHERE id=$1`, productsTable)
	if err := p.db.QueryRow(ctx, query, input.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
	); err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, errors.New("product not found")
		}
		return nil, response.NewRpcErrResponse(codes.NotFound, errors.New(err))
	}
	resp := &v1.GetProductByIdResponse{Product: &product}
	return resp, nil
}
func (p *ProductRepo) DeleteProduct(ctx context.Context, input *v1.DeleteProductRequest) (*v1.DeleteProductResponse, error) {
	var resp *v1.DeleteProductResponse
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, productsTable)
	exec, err := p.db.Exec(ctx, query, input.ProductId)
	if err != nil {
		return nil, response.NewRpcErrResponse(codes.Internal, errors.New(err))
	}
	if rowsAffected := exec.RowsAffected(); rowsAffected > 0 {
		resp = &v1.DeleteProductResponse{
			StatusCode: http.StatusOK,
			Message:    "success",
		}
		return resp, nil
	}
	resp = &v1.DeleteProductResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "error",
	}
	return resp, nil
}

func UpdateProduct() {

}
