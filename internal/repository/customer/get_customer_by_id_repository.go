package customerRepository

import (
	"context"
	"database/sql"
	"eniqlo/internal/entity"
	"fmt"
)

func (cr customerRepository) GetCustomerByID(ctx context.Context, id string) (*entity.Customer, error) {

	var (
		resp entity.Customer
		err  error
	)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE "id" = $1`, resp.TableName())

	err = cr.db.GetContext(ctx, &resp, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &resp, nil
}
