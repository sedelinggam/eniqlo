package productController

import (
	productHandler "eniqlo/internal/delivery/http/v1/controller/product/handler"
	productService "eniqlo/internal/service/product"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, productSvc productService.ProductService) {
	user := group.Group("/product")
	handler := productHandler.NewHandler(productSvc, val)

	publicRoute := user
	publicRoute.Use()
	publicRoute.POST("", handler.CreateProduct)
	publicRoute.PUT("/:id", handler.CreateProduct)
}
