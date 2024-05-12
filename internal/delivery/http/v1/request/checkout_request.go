package request

type GetCheckoutHistoriesFilterKey string

const (
	ShouldGetCheckoutHistoriesFilterKey GetCheckoutHistoriesFilterKey = "getCheckoutHistoriesFilter"
)

type ShouldGetCheckoutHistoriesFilter struct {
	CustomerID bool
	Limit      bool
	Offset     bool
	CreatedAt  bool
}

type GetCheckoutHistories struct {
	CustomerID string
	Limit      int32
	Offset     int32
	CreatedAt  string
}

type CheckoutProduct struct {
	CustomerID     string `json:"customerId" validate:"required"`
	ProductDetails []struct {
		ProductID string `json:"productId" validate:"required"`
		Quantity  uint   `json:"quantity" validate:"required,gte=1"`
	} `json:"productDetails"`
	Paid   int `json:"paid" validate:"required,gte=0"`
	Change int `json:"change" validate:"required,gte=0"`
}
