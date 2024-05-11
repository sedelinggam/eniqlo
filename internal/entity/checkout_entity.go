package entity

import "time"

type Checkout struct {
	ID         string    `db:"id"`
	CustomerID string    `db:"customer_id"`
	Paid       int       `db:"paid"`
	Change     int       `db:"change"`
	CreatedAt  time.Time `db:"created_at"`
}

func (c Checkout) TableName() string {
	return `checkouts`
}
