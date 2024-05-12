package checkoutRepository

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/entity"
	"fmt"
	"strings"
)

func (cr checkoutRepository) GetCheckoutHistories(ctx context.Context, req request.GetCheckoutHistories) (*[]entity.Checkout, error) {
	var (
		conditions []string
		filter     []interface{}
		resp       []entity.Checkout
		err        error
	)

	shouldFilter := ctx.Value(request.ShouldGetCheckoutHistoriesFilterKey).(request.ShouldGetCheckoutHistoriesFilter)

	query := fmt.Sprintf(`SELECT c.id, c.customer_id, cd.product_id, cd.quantity, c.paid, c.change, c.created_at
    FROM %s c
    JOIN %s cd ON c.id = cd.checkout_id`, entity.Checkout{}.TableName(), entity.CheckoutDetail{}.TableName())

	if shouldFilter.CustomerID {
		filter = append(filter, req.CustomerID)
		conditions = append(conditions, fmt.Sprintf("c.customer_id = $%d", len(filter)))
	}

	if len(conditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
	}

	if shouldFilter.CreatedAt {
		fmt.Println(shouldFilter.CreatedAt, req.CreatedAt, "AAAAAAAAAAAAAAAAAAAAA")
		query += fmt.Sprintf(" ORDER BY c.created_at %s", req.CreatedAt)
	} else {
		query += " ORDER BY c.created_at DESC"
	}

	filter = append(filter, req.Limit)
	query += fmt.Sprintf(" LIMIT $%d", len(filter))

	filter = append(filter, req.Offset)
	fmt.Println(req.Offset, "AAAAAAAAAAAAAAAAAAAAA")
	query += fmt.Sprintf(" OFFSET $%d", len(filter))

	rows := []entity.CheckoutRow{}
	err = cr.db.Select(&rows, query, filter...)
	if err != nil {
		return nil, err
	}

	checkouts := map[string]*entity.Checkout{}
	for _, row := range rows {
		checkout, exists := checkouts[row.ID]
		if !exists {
			checkout = &entity.Checkout{
				ID:         row.ID,
				CustomerID: row.CustomerID,
				Paid:       row.Paid,
				Change:     row.Change,
				CreatedAt:  row.CreatedAt,
			}
			checkouts[row.ID] = checkout
		}
		checkout.CheckoutDetails = append(checkout.CheckoutDetails, &entity.CheckoutDetail{
			ProductID: row.ProductID,
			Quantity:  uint(row.Quantity),
		})
	}

	resp = []entity.Checkout{}
	for _, checkout := range checkouts {
		resp = append(resp, *checkout)
	}

	return &resp, nil
}
