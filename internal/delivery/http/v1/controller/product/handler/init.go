package productHandler

import (
	checkoutService "eniqlo/internal/service/checkout"
	productService "eniqlo/internal/service/product"

	"github.com/go-playground/validator/v10"
)

type productHandler struct {
	productService  productService.ProductService
	checkoutService checkoutService.CheckoutService
	val             *validator.Validate
}

func NewHandler(
	productService productService.ProductService,
	checkoutService checkoutService.CheckoutService,
	val *validator.Validate) *productHandler {
	return &productHandler{
		productService:  productService,
		checkoutService: checkoutService,
		val:             val,
	}
}
