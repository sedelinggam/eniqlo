package productService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/internal/entity"
	"eniqlo/package/lumen"
	"time"

	"github.com/lib/pq"
)

func (ps productService) DeleteProducts(ctx context.Context, requestData request.DeleteProduct) (*response.DeleteProduct, error) {
	//Password Hash
	var (
		err error
	)

	//Update Cat
	catData := entity.Product{
		ID: requestData.ID,
		DeletedAt: pq.NullTime{
			Time: time.Now(),
		},
	}

	err = ps.productRepo.UpdateDeletedAt(ctx, catData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.DeleteProduct{
		ID:            catData.ID,
		DeleteProduct: time.Now().Format(time.RFC3339),
	}, nil
}
