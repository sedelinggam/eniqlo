package customerRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (cr customerRepository) Register(ctx context.Context, data entity.Customer) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, phone_number, name, created_at) VALUES (:id, :phone_number, :name, :created_at)`, data.TableName())

	tx := cr.db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
