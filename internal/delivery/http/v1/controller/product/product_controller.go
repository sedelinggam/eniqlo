package productController

import (
	productHandler "eniqlo/internal/delivery/http/v1/controller/product/handler"
	productService "eniqlo/internal/service/product"
	cryptoJWT "eniqlo/package/crypto/jwt"

	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, productSvc productService.ProductService) {
	user := group.Group("/product")
	handler := productHandler.NewHandler(productSvc, val)

	privateRoute := user
	privateRoute.Use(echojwt.WithConfig(cryptoJWT.JWTConfig()))
	privateRoute.POST("", handler.CreateProduct)
	privateRoute.PUT("/:id", handler.UpdateProduct)
	privateRoute.DELETE("/:id", handler.DeleteProduct)
	privateRoute.GET("", handler.GetProducts)
}
