package productRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (cr productRepository) Create(ctx context.Context, data entity.Product) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, name, sku, category, image_url, notes, price, stock, location, is_avaiable, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`, data.TableName())

	tx := cr.db.MustBegin()
	_, err := tx.Exec(query, data.ID, data.Name, data.SKU, data.Category, data.ImageUrl, data.Notes, data.Price, data.Stock, data.Location, data.IsAvailable, data.CreatedAt)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}
