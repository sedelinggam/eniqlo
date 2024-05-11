package productService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/internal/entity"
	valueobject "eniqlo/internal/value_object"
	"eniqlo/package/lumen"
	"time"

	"github.com/google/uuid"
)

func (ps productService) CreateProduct(ctx context.Context, requestData request.CreateProduct) (*response.CreateProduct, error) {

	var (
		err error
	)

	//Check product category
	err = valueobject.CheckProductCategory(requestData.Category)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	//Create Cat
	catData := entity.Product{
		ID:          uuid.New().String(),
		Name:        requestData.Name,
		SKU:         requestData.Sku,
		Category:    requestData.Category,
		ImageUrl:    requestData.ImageURL,
		Notes:       requestData.Notes,
		Price:       requestData.Price,
		Stock:       *requestData.Stock,
		Location:    requestData.Location,
		IsAvailable: *requestData.IsAvailable,
		CreatedAt:   time.Now(),
	}

	err = ps.productRepo.Create(ctx, catData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.CreateProduct{
		ID:        catData.ID,
		CreatedAt: catData.CreatedAt.Format(time.RFC3339),
	}, nil
}
