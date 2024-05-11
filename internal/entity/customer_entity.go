package entity

import (
	"eniqlo/package/lumen"
	"errors"
	"strings"
	"time"
)

type Customer struct {
	ID          string    `db:"id"`
	PhoneNumber string    `db:"phone_number"`
	Name        string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
}

func (c Customer) CheckPhoneNumber() error {
	if len(c.PhoneNumber) == 0 {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("phone number not valid"))
	}
	if !strings.HasPrefix(c.PhoneNumber, "+") {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("phone number not valid"))
	}
	return nil
}

func (c Customer) TableName() string {
	return `customers`
}
