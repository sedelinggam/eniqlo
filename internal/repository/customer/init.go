package customerRepository

import (
	"context"
	"eniqlo/internal/entity"

	"github.com/jmoiron/sqlx"
)

type customerRepository struct {
	db *sqlx.DB
}

type CustomerRepository interface {
	Register(ctx context.Context, data entity.Customer) error
	GetCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.Customer, error)
	GetCustomerByID(ctx context.Context, id string) (*entity.Customer, error)
}

func New(db *sqlx.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
