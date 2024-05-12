package checkoutService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/package/lumen"
)

func (cs checkoutService) GetCheckoutHistories(ctx context.Context, requestData request.GetCheckoutHistories) (*[]response.GetCheckoutHistories, error) {
	resp := &[]response.GetCheckoutHistories{}

	checkouts, err := cs.checkoutRepo.GetCheckoutHistories(ctx, requestData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	for _, checkout := range *checkouts {
		history := response.GetCheckoutHistories{
			TransactionID: checkout.ID,
			CustomerID:    checkout.CustomerID,
			Paid:          checkout.Paid,
			Change:        checkout.Change,
			CreatedAt:     checkout.CreatedAt,
		}

		for _, detail := range checkout.CheckoutDetails {
			history.ProductDetails = append(history.ProductDetails, struct {
				ProductID string `json:"productId"`
				Quantity  uint   `json:"quantity"`
			}{
				ProductID: detail.ProductID,
				Quantity:  detail.Quantity,
			})
		}

		*resp = append(*resp, history)
	}

	return resp, nil
}
