package response

import "time"

type CheckoutResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
}

type GetCheckoutHistories struct {
	TransactionID  string `json:"transactionId"`
	CustomerID     string `json:"customerId"`
	ProductDetails []struct {
		ProductID string `json:"productId"`
		Quantity  uint   `json:"quantity"`
	} `json:"productDetails"`
	Paid      int       `json:"paid"`
	Change    int       `json:"change"`
	CreatedAt time.Time `json:"createdAt"`
}
