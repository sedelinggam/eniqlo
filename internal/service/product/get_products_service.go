package productService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/package/lumen"
)

func (ps productService) GetProducts(ctx context.Context, req request.GetProducts) (*[]response.GetProducts, error) {
	products, err := ps.productRepo.Gets(ctx, req)

	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}

		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	productsSlice := *products
	productsResp := []response.GetProducts{}

	for i := 0; i < len(productsSlice); i++ {
		productRaw := productsSlice[i]
		product := response.GetProducts{
			ID:          productRaw.ID,
			Name:        productRaw.Name,
			Sku:         productRaw.SKU,
			Category:    productRaw.Category,
			Stock:       productRaw.Stock,
			Notes:       productRaw.Notes,
			Price:       productRaw.Price,
			Location:    productRaw.Location,
			IsAvailable: productRaw.IsAvailable,
			CreatedAt:   productRaw.CreatedAt,
			ImageUrl:    productRaw.ImageUrl,
		}
		productsResp = append(productsResp, product)
	}

	return &productsResp, nil
}
