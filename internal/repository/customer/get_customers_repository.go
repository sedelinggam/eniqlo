package customerRepository

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/entity"
	"fmt"
	"strings"
)

func (pr customerRepository) GetCustomers(ctx context.Context, req request.Customer) ([]*entity.Customer, error) {
	var (
		conditions []string
		filter     []interface{}
		resp       []*entity.Customer
		err        error
		model      entity.Customer
	)

	query := fmt.Sprintf(`SELECT * FROM %s`, model.TableName())

	if req.Name != nil {
		filter = append(filter, req.Name)
		conditions = append(conditions, fmt.Sprintf("name ILIKE '%%' || $%d || '%%'", len(filter)))
	}

	if req.PhoneNumber != nil {
		filter = append(filter, req.PhoneNumber)
		conditions = append(conditions, fmt.Sprintf("phone_number ILIKE $%d || '%%'", len(filter)))
	}

	if len(conditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
	}

	query += " ORDER BY created_at ASC"

	fmt.Println(query)
	err = pr.db.Select(&resp, query, filter...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
