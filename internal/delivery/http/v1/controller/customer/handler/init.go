package customerHandler

import (
	customerService "eniqlo/internal/service/customer"

	"github.com/go-playground/validator/v10"
)

type customerHandler struct {
	customerService customerService.CustomerService
	val             *validator.Validate
}

func NewHandler(customerService customerService.CustomerService, val *validator.Validate) *customerHandler {
	return &customerHandler{
		customerService: customerService,
		val:             val,
	}
}
