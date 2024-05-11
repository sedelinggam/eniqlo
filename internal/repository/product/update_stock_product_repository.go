package productRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (pr productRepository) UpdateStock(ctx context.Context, id string, stock int) error {

	var product entity.Product
	var query string

	if stock <= 0 {
		query = fmt.Sprintf(`UPDATE %s SET stock = 0, is_available = false WHERE id = $1`, product.TableName())
		_, err := pr.db.ExecContext(ctx, query, id)
		if err != nil {
			return err
		}
	} else {
		query = fmt.Sprintf(`UPDATE %s SET stock = $1 WHERE id = $2`, product.TableName())
		_, err := pr.db.ExecContext(ctx, query, stock, id)
		if err != nil {
			return err
		}
	}

	return nil
}
