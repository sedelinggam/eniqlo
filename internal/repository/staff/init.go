package staffRepository

import (
	"context"
	"eniqlo/internal/entity"

	"github.com/jmoiron/sqlx"
)

type staffRepository struct {
	db *sqlx.DB
}

type StaffRepository interface {
	Create(ctx context.Context, data entity.Staff) error
	GetStaffByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.Staff, error)
}

func New(db *sqlx.DB) StaffRepository {
	return &staffRepository{
		db: db,
	}
}
