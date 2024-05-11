package entity

type CheckoutDetail struct {
	ID         string `db:"id"`
	CheckoutID string `db:"checkout_id"`
	ProductID  string `db:"product_id"`
	Quantity   uint   `db:"quantity"`
}

func (c CheckoutDetail) TableName() string {
	return `checkout_details`
}
