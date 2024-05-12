package checkoutDetailRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (cr checkoutDetailRepository) CreateCheckoutDetail(ctx context.Context, data entity.CheckoutDetail) error {

	query := fmt.Sprintf(`INSERT INTO %s(id, product_id, checkout_id, quantity) VALUES (:id, :product_id, :checkout_id, :quantity)`, data.TableName())

	tx := cr.db.MustBegin()
	// Execute the query without starting a transaction
	_, err := tx.NamedExecContext(ctx, query, data)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
