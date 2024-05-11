package customerService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	customerRepository "eniqlo/internal/repository/customer"

	"github.com/jmoiron/sqlx"
)

type customerService struct {
	customerRepo customerRepository.CustomerRepository
}

type CustomerService interface {
	Register(ctx context.Context, requestData request.CustomerRegister) (*response.CustomerRegister, error)
}

func New(db *sqlx.DB) CustomerService {
	return &customerService{
		customerRepo: customerRepository.New(db),
	}
}
