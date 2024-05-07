package staffService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	staffRepository "eniqlo/internal/repository/staff"

	"github.com/jmoiron/sqlx"
)

type staffService struct {
	staffRepo staffRepository.StaffRepository
}

type StaffService interface {
	Login(ctx context.Context, requestData request.StaffLogin) (*response.UserAccessToken, error)
	Register(ctx context.Context, requestData request.StaffRegister) (*response.UserAccessToken, error)
}

func New(db *sqlx.DB) StaffService {
	return &staffService{
		staffRepo: staffRepository.New(db),
	}
}
