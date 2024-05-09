package productRepository

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/entity"
	"fmt"
	"strings"
)

func (pr productRepository) Gets(ctx context.Context, req request.GetProducts) (*[]entity.Product, error) {
	var (
		conditions []string
		filter     []interface{}
		resp       []entity.Product
		err        error
	)

	shouldFilter := ctx.Value(request.ShouldGetProductsFilterKey).(request.ShouldGetProductsFilter)

	query := `SELECT id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at FROM products`

	if shouldFilter.ID {
		filter = append(filter, req.ID)
		conditions = append(conditions, fmt.Sprintf("id = $%d", len(filter)))
	}

	if shouldFilter.Name {
		filter = append(filter, req.Name)
		conditions = append(conditions, fmt.Sprintf("name ILIKE '%%' || $%d || '%%'", len(filter)))
	}

	if shouldFilter.Category {
		filter = append(filter, req.Category)
		conditions = append(conditions, fmt.Sprintf("category = $%d", len(filter)))
	}

	if shouldFilter.IsAvailable {
		filter = append(filter, req.IsAvailable)
		conditions = append(conditions, fmt.Sprintf("is_available = $%d", len(filter)))
	}

	if shouldFilter.Sku {
		filter = append(filter, req.Sku)
		conditions = append(conditions, fmt.Sprintf("sku = $%d", len(filter)))
	}

	if shouldFilter.InStock {
		if req.InStock {
			conditions = append(conditions, "stock > 0")
		} else {
			conditions = append(conditions, "stock = 0")
		}
	}

	if len(conditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
		query += " AND deleted_at IS NULL"
	} else {
		query += " WHERE deleted_at IS NULL"
	}

	if shouldFilter.Price && shouldFilter.CreatedAt {
		query += fmt.Sprintf(" ORDER BY price %s, created_at %s", req.Price, req.CreatedAt)
	} else if shouldFilter.Price {
		query += fmt.Sprintf(" ORDER BY price %s", req.Price)
	} else if shouldFilter.CreatedAt {
		query += fmt.Sprintf(" ORDER BY created_at %s", req.CreatedAt)
	} else {
		query += " ORDER BY created_at ASC"
	}

	filter = append(filter, req.Limit)
	query += fmt.Sprintf(" LIMIT $%d", len(filter))

	filter = append(filter, req.Offset)
	query += fmt.Sprintf(" OFFSET $%d", len(filter))

	fmt.Println(query)

	err = pr.db.Select(&resp, query, filter...)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
