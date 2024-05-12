package customerService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/package/lumen"
)

func (cs customerService) GetCustomers(ctx context.Context, requestData request.Customer) ([]*response.Customer, error) {
	customer, err := cs.customerRepo.GetCustomers(ctx, requestData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return response.MapCustomerListEntityToListResponse(customer), nil
}
