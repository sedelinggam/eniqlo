package customerController

import (
	customerHandler "eniqlo/internal/delivery/http/v1/controller/customer/handler"
	customerService "eniqlo/internal/service/customer"
	cryptoJWT "eniqlo/package/crypto/jwt"

	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, customerSvc customerService.CustomerService) {
	user := group.Group("/customer")
	handler := customerHandler.NewHandler(customerSvc, val)

	privateRoute := user
	privateRoute.Use(echojwt.WithConfig(cryptoJWT.JWTConfig()))

	privateRoute.POST("/register", handler.RegisterCustomer)

}
