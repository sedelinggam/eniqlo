package staffControllers

import (
	staffHandler "eniqlo/internal/delivery/http/v1/controller/staff/handler"
	staffService "eniqlo/internal/service/staff"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, staffSvc staffService.StaffService) {
	user := group.Group("/staff")
	handler := staffHandler.NewHandler(staffSvc, val)

	publicRoute := user
	publicRoute.Use()
	publicRoute.POST("/register", handler.Register)
	publicRoute.POST("/login", handler.Login)
}
