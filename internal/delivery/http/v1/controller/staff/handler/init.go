package staffHandler

import (
	staffService "eniqlo/internal/service/staff"

	"github.com/go-playground/validator/v10"
)

type staffHandler struct {
	staffService staffService.StaffService
	val          *validator.Validate
}

func NewHandler(staffService staffService.StaffService, val *validator.Validate) *staffHandler {
	return &staffHandler{
		staffService: staffService,
		val:          val,
	}
}
