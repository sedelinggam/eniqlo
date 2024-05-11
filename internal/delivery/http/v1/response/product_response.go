package response

import "time"

type CreateProduct struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type UpdateProduct struct {
	ID        string `json:"id"`
	UpdatedAt string `json:"updatedAt"`
}

type DeleteProduct struct {
	ID            string `json:"id"`
	DeleteProduct string `json:"deletedAt"`
}

type GetProducts struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Sku         string    `json:"sku"`
	Category    string    `json:"category"`
	ImageUrl    string    `json:"imageUrl"`
	Stock       uint      `json:"stock"`
	Notes       string    `json:"notes"`
	Price       uint      `json:"price"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"isAvailable"`
	CreatedAt   time.Time `json:"createdAt"`
}
