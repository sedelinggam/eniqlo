package checkoutService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/internal/entity"
	"eniqlo/package/lumen"
	"errors"
	"time"

	"github.com/google/uuid"
)

func (cs checkoutService) CheckoutProduct(ctx context.Context, requestData request.CheckoutProduct) (*response.CheckoutResponse, error) {
	totalPrice := uint(0)

	// get product detail
	for _, productReq := range requestData.ProductDetails {
		product, err := cs.productRepo.GetProductByID(ctx, productReq.ProductID)
		// check if product not found
		if err != nil {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}

		// check if product is not available
		if !product.IsAvailable {
			return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("product is not available"))
		}

		// check stock
		if product.Stock < uint(productReq.Quantity) {
			return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("stock is not enough"))
		}

		// calculate total price
		totalPrice += product.Price * uint(productReq.Quantity)
	}

	// validate paid amount
	if uint(requestData.Paid) < totalPrice {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("paid amount is not enough"))
	}

	// validate change amount
	changeAmount := uint(requestData.Paid) - totalPrice
	if changeAmount != uint(requestData.Change) {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("change amount is not valid"))

	}

	//Create checkout
	checkout := entity.Checkout{
		ID:         uuid.New().String(),
		CustomerID: requestData.CustomerID,
		Paid:       requestData.Paid,
		Change:     requestData.Change,
		CreatedAt:  time.Time{},
	}

	err := cs.checkoutRepo.CreateCheckout(ctx, checkout)

	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Update product stock
	for _, productReq := range requestData.ProductDetails {
		product, _ := cs.productRepo.GetProductByID(ctx, productReq.ProductID)
		newStock := product.Stock - uint(productReq.Quantity)
		err = cs.productRepo.UpdateStock(ctx, productReq.ProductID, int(newStock))
		if err != nil {
			return nil, lumen.NewError(lumen.ErrInternalFailure, err)
		}
	}

	return &response.CheckoutResponse{
		ID:         checkout.ID,
		CustomerID: checkout.CustomerID,
	}, nil
}
