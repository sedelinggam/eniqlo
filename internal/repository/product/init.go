package productRepository

import (
	"context"
	"eniqlo/internal/entity"

	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

type ProductRepository interface {
	Create(ctx context.Context, data entity.Product) error
}

func New(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
