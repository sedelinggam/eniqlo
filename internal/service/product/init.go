package productService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	productRepository "eniqlo/internal/repository/product"

	"github.com/jmoiron/sqlx"
)

type productService struct {
	productRepo productRepository.ProductRepository
}

type ProductService interface {
	CreateProduct(ctx context.Context, requestData request.CreateProduct) (*response.CreateProduct, error)
	UpdateProduct(ctx context.Context, requestData request.UpdateProduct) (*response.UpdateProduct, error)
	GetProducts(ctx context.Context, requestData request.GetProducts) (*[]response.GetProducts, error)
	DeleteProducts(ctx context.Context, requestData request.DeleteProduct) (*response.DeleteProduct, error)
}

func New(db *sqlx.DB) ProductService {
	return &productService{
		productRepo: productRepository.New(db),
	}
}
