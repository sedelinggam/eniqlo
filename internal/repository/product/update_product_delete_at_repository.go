package productRepository

import (
	"context"
	"eniqlo/internal/entity"
	"errors"
	"fmt"
)

func (cr productRepository) UpdateDeletedAt(ctx context.Context, data entity.Product) error {
	query := fmt.Sprintf(`UPDATE %s SET deleted_at = $1 WHERE id = $2`, data.TableName())

	tx := cr.db.MustBegin()
	res, err := tx.Exec(query, data.DeletedAt.Time, data.ID)
	tx.Commit()
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("no rows in result set")
	}

	return nil
}
