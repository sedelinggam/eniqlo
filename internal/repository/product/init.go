package productRepository

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/entity"

	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

type ProductRepository interface {
	Create(ctx context.Context, data entity.Product) error
	Update(ctx context.Context, data entity.Product) error
	Gets(ctx context.Context, req request.GetProducts) (*[]entity.Product, error)
}

func New(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
