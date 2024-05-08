package request

type CreateProduct struct {
	Name        string `json:"name" validate:"required,min=1,max=30"`
	Sku         string `json:"sku" validate:"required,min=1,max=30"`
	Category    string `json:"category" validate:"required"`
	ImageURL    string `json:"imageUrl" validate:"required"`
	Notes       string `json:"notes" validate:"required,min=1,max=200"`
	Price       uint   `json:"price" validate:"required,gte=1"`
	Stock       uint   `json:"stock" validate:"required,gte=0,lte=100000"`
	Location    string `json:"location" validate:"required,min=1,max=200"`
	IsAvailable bool   `json:"isAvailable" validate:"required"`
}
