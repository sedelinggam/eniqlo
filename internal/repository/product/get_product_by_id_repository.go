package productRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (pr productRepository) GetProductByID(ctx context.Context, id string) (entity.Product, error) {
	var product entity.Product

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1 AND "deleted_at" IS NULL`, product.TableName())

	err := pr.db.GetContext(ctx, &product, query, id)
	if err != nil {
		return product, err
	}

	return product, nil
}
