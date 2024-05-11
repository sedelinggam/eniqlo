package checkoutService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	checkoutRepository "eniqlo/internal/repository/checkout"
	checkoutDetailRepository "eniqlo/internal/repository/checkout_detail"
	customerRepository "eniqlo/internal/repository/customer"
	productRepository "eniqlo/internal/repository/product"

	"github.com/jmoiron/sqlx"
)

type checkoutService struct {
	customerRepo       customerRepository.CustomerRepository
	productRepo        productRepository.ProductRepository
	checkoutRepo       checkoutRepository.CheckoutRepository
	checkoutDetailRepo checkoutDetailRepository.CheckoutDetailRepository
}

type CheckoutService interface {
	CheckoutProduct(ctx context.Context, requestData request.CheckoutProduct) (*response.CheckoutResponse, error)
}

func New(db *sqlx.DB) CheckoutService {
	return &checkoutService{
		customerRepo:       customerRepository.New(db),
		productRepo:        productRepository.New(db),
		checkoutRepo:       checkoutRepository.New(db),
		checkoutDetailRepo: checkoutDetailRepository.New(db),
	}
}
