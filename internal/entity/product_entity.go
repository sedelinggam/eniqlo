package entity

import (
	"time"
)

type Product struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	SKU         string    `db:"sku"`
	Category    string    `db:"category"`
	ImageUrl    string    `db:"image_url"`
	Notes       string    `db:"notes"`
	Price       uint      `db:"price"`
	Stock       uint      `db:"stock"`
	Location    string    `db:"location"`
	IsAvailable bool      `db:"is_available"`
	CreatedAt   time.Time `db:"created_at"`
}

func (p Product) TableName() string {
	return `products`
}
