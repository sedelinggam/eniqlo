package productRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (cr productRepository) UpdateDeletedAt(ctx context.Context, data entity.Product) error {
	query := fmt.Sprintf(`UPDATE %s SET deleted_at = $1 WHERE id = $2`, data.TableName())

	tx := cr.db.MustBegin()
	_, err := tx.Exec(query, data.DeletedAt.Time, data.ID)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}
