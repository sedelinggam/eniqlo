package checkoutService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	checkoutRepository "eniqlo/internal/repository/checkout"
	productRepository "eniqlo/internal/repository/product"

	"github.com/jmoiron/sqlx"
)

type checkoutService struct {
	productRepo  productRepository.ProductRepository
	checkoutRepo checkoutRepository.CheckoutRepository
}

type CheckoutService interface {
	CheckoutProduct(ctx context.Context, requestData request.CheckoutProduct) (*response.CheckoutResponse, error)
}

func New(db *sqlx.DB) CheckoutService {
	return &checkoutService{
		productRepo:  productRepository.New(db),
		checkoutRepo: checkoutRepository.New(db),
	}
}
