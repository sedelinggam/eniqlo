package customerService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/internal/entity"
	"eniqlo/package/lumen"
	"time"

	"github.com/google/uuid"
)

func (cs customerService) Register(ctx context.Context, requestData request.CustomerRegister) (*response.CustomerRegister, error) {

	existCustomer, err := cs.customerRepo.GetCustomerByPhoneNumber(ctx, requestData.PhoneNumber)

	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	if existCustomer != nil {
		return nil, lumen.NewError(lumen.ErrConflict, lumen.ErrConflict)
	}

	customer := entity.Customer{
		ID:          uuid.New().String(),
		PhoneNumber: requestData.PhoneNumber,
		Name:        requestData.Name,
		CreatedAt:   time.Now(),
	}

	err = customer.CheckPhoneNumber()
	if err != nil {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	err = cs.customerRepo.Register(ctx, customer)
	if err != nil {
		return nil, err
	}

	return &response.CustomerRegister{
		UserID:      customer.ID,
		PhoneNumber: customer.PhoneNumber,
		Name:        customer.Name,
	}, nil
}
