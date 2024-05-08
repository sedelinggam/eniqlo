package staffRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (sr staffRepository) Create(ctx context.Context, data entity.Staff) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, phone_number, name, password, created_at) VALUES (:id, :phone_number, :name, :password, :created_at)`, data.TableName())

	tx := sr.db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
