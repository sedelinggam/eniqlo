package checkoutRepository

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/entity"

	"github.com/jmoiron/sqlx"
)

type checkoutRepository struct {
	db *sqlx.DB
}

type CheckoutRepository interface {
	CreateCheckout(ctx context.Context, data entity.Checkout) error
	GetCheckoutHistories(ctx context.Context, req request.GetCheckoutHistories) (*[]entity.Checkout, error)
}

func New(db *sqlx.DB) CheckoutRepository {
	return &checkoutRepository{
		db: db,
	}
}
