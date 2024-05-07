package entity

import "time"

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

func (s Staff) NewPhoneNumber(phoneNumber string) (string, error) {
	return "", nil
}

func (s Staff) NewPassword(password string) (string, error) {
	return "", nil
}
