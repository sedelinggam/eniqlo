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

type UpdateProduct struct {
	ID          string `json:"id"`
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

type GetProducts struct {
	ID          string
	Limit       int32
	Offset      int32
	Name        string
	IsAvailable bool
	Category    string
	Sku         string
	Price       string
	InStock     bool
	CreatedAt   string
}

type DeleteProduct struct {
	ID string
}

type GetProductsFilterKey string

const (
	ShouldGetProductsFilterKey GetProductsFilterKey = "getProductsFilter"
)

type ShouldGetProductsFilter struct {
	ID          bool
	Limit       bool
	Offset      bool
	Name        bool
	IsAvailable bool
	Category    bool
	Sku         bool
	Price       bool
	InStock     bool
	CreatedAt   bool
}
