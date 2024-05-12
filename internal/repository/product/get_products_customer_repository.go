package productRepository

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/entity"
	"fmt"
	"strings"
)

func (pr productRepository) GetCustomerProducts(ctx context.Context, req request.GetCustomerProducts) (*[]entity.Product, error) {
	var (
		conditions []string
		filter     []interface{}
		resp       []entity.Product
		err        error
	)

	query := `SELECT id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at FROM products`

	if req.Name != nil {
		filter = append(filter, req.Name)
		conditions = append(conditions, fmt.Sprintf("name ILIKE '%%' || $%d || '%%'", len(filter)))
	}

	if req.Category != nil {
		filter = append(filter, req.Category)
		conditions = append(conditions, fmt.Sprintf("category = $%d", len(filter)))
	}

	if req.Sku != nil {
		filter = append(filter, req.Sku)
		conditions = append(conditions, fmt.Sprintf("sku = $%d", len(filter)))
	}

	if req.InStock != nil {
		if *req.InStock {
			conditions = append(conditions, "stock > 0")
		} else {
			conditions = append(conditions, "stock = 0")
		}
	}

	if len(conditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
		query += " AND is_available = true AND deleted_at IS NULL"
	} else {
		query += " WHERE is_available = true AND deleted_at IS NULL"
	}

	if req.Price != nil && req.CreatedAt != nil {
		query += fmt.Sprintf(" ORDER BY price %s, created_at %s", *req.Price, *req.CreatedAt)
	} else if req.Price != nil {
		query += fmt.Sprintf(" ORDER BY price %s", *req.Price)
	} else if req.CreatedAt != nil {
		query += fmt.Sprintf(" ORDER BY created_at %s", *req.CreatedAt)
	} else {
		query += " ORDER BY created_at DESC"
	}

	filter = append(filter, req.Limit)
	query += fmt.Sprintf(" LIMIT $%d", len(filter))

	filter = append(filter, req.Offset)
	query += fmt.Sprintf(" OFFSET $%d", len(filter))

	err = pr.db.Select(&resp, query, filter...)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
