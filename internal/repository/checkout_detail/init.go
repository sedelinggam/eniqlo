package checkoutDetailRepository

import (
	"context"
	"eniqlo/internal/entity"

	"github.com/jmoiron/sqlx"
)

type checkoutDetailRepository struct {
	db *sqlx.DB
}

type CheckoutDetailRepository interface {
	CreateCheckoutDetail(ctx context.Context, data entity.CheckoutDetail) error
}

func New(db *sqlx.DB) CheckoutDetailRepository {
	return &checkoutDetailRepository{
		db: db,
	}
}
