package entity

import (
	"eniqlo/package/lumen"
	"errors"
	"strings"
	"time"
)

type Staff struct {
	ID          string    `db:"id"`
	PhoneNumber string    `db:"phone_number"`
	Name        string    `db:"name"`
	Password    string    `db:"password"`
	CreatedAt   time.Time `db:"created_at"`
}

func (s Staff) TableName() string {
	return `staffs`
}

func (s Staff) CheckPhoneNumber() error {
	if len(s.PhoneNumber) == 0 {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("phone number not valid"))
	}
	if !strings.HasPrefix(s.PhoneNumber, "+") {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("phone number not valid"))
	}
	return nil
}

func (s *Staff) NewPassword(password string) error {
	return nil
}
