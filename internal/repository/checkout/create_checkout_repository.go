package checkoutRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (cr checkoutRepository) CreateCheckout(ctx context.Context, data entity.Checkout) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, customer_id, paid, change, created_at) VALUES (:id, :customer_id, :paid, :change, :created_at)`, data.TableName())

	// Execute the query without starting a transaction
	_, err := cr.db.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}

	return nil
}
