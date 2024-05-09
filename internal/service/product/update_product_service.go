package productService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/internal/entity"
	"eniqlo/package/lumen"
	"time"
)

func (ps productService) UpdateProduct(ctx context.Context, requestData request.UpdateProduct) (*response.UpdateProduct, error) {
	//Password Hash
	var (
		err error
	)

	//Update Cat
	catData := entity.Product{
		ID:          requestData.ID,
		Name:        requestData.Name,
		SKU:         requestData.Sku,
		Category:    requestData.Category,
		ImageUrl:    requestData.ImageURL,
		Notes:       requestData.Notes,
		Price:       requestData.Price,
		Stock:       requestData.Stock,
		Location:    requestData.Location,
		IsAvailable: requestData.IsAvailable,
	}

	err = ps.productRepo.Update(ctx, catData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.UpdateProduct{
		ID:        catData.ID,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil
}
