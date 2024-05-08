package productHandler

import (
	productService "eniqlo/internal/service/product"

	"github.com/go-playground/validator/v10"
)

type productHandler struct {
	productService productService.ProductService
	val            *validator.Validate
}

func NewHandler(productService productService.ProductService, val *validator.Validate) *productHandler {
	return &productHandler{
		productService: productService,
		val:            val,
	}
}
