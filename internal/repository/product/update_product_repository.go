package productRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (cr productRepository) Update(ctx context.Context, data entity.Product) error {
	query := fmt.Sprintf(`UPDATE %s SET(name, sku, category, image_url, notes, price, stock, location, is_available) = ($1,$2,$3,$4,$5,$6,$7,$8,$9) WHERE id = $10`, data.TableName())

	tx := cr.db.MustBegin()
	_, err := tx.Exec(query, data.Name, data.SKU, data.Category, data.ImageUrl, data.Notes, data.Price, data.Stock, data.Location, data.IsAvailable, data.ID)
	tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
