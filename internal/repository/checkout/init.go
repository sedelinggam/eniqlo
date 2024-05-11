package checkoutRepository

import (
	"context"
	"eniqlo/internal/entity"

	"github.com/jmoiron/sqlx"
)

type checkoutRepository struct {
	db *sqlx.DB
}

type CheckoutRepository interface {
	CreateCheckout(ctx context.Context, data entity.Checkout) error
}

func New(db *sqlx.DB) CheckoutRepository {
	return &checkoutRepository{
		db: db,
	}
}
