package productHandler

import (
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/package/lumen"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ph productHandler) GetCustomerProducts(c echo.Context) error {
	var (
		req  request.GetCustomerProducts
		resp *[]response.GetProducts
		err  error
	)

	queries := c.QueryParams()

	if name := queries.Get("name"); name != "" {
		name := queries.Get("name")
		req.Name = &name
	}

	if sku := queries.Get("sku"); sku != "" {
		sku := queries.Get("sku")
		req.Sku = &sku
	}

	if category := queries.Get("category"); category != "" {
		err := ph.val.Var(queries.Get("category"), "oneof=Clothing Accessories Footwear Beverages")
		if err == nil {
			category := queries.Get("category")
			req.Category = &category
		}
	}

	if price := queries.Get("price"); price != "" {
		err := ph.val.Var(queries.Get("price"), "oneof=asc desc")
		if err == nil {
			price := queries.Get("price")
			req.Price = &price
		}
	}

	if inStock := queries.Get("inStock"); inStock != "" {
		err := ph.val.Var(queries.Get("inStock"), "boolean")
		if err == nil {
			val, _ := strconv.ParseBool(queries.Get("inStock"))
			req.InStock = &val
		}
	}

	if createdAt := queries.Get("createdAt"); createdAt != "" {
		err := ph.val.Var(queries.Get("createdAt"), "oneof=asc desc")
		if err == nil {
			createdAt := queries.Get("createdAt")
			req.CreatedAt = &createdAt
		}
	}

	if limit := queries.Get("limit"); limit != "" {
		err := ph.val.Var(queries.Get("limit"), "number")
		if err == nil {
			val, _ := strconv.ParseInt(queries.Get("limit"), 10, 32)
			req.Limit = int32(val)
		}
	} else {
		req.Limit = 5
	}

	if offset := queries.Get("offset"); offset != "" {
		err := ph.val.Var(queries.Get("offset"), "number")
		if err == nil {
			val, _ := strconv.ParseInt(queries.Get("offset"), 10, 32)
			req.Offset = int32(val)
		}
	} else {
		req.Offset = 0
	}

	ctx := c.Request().Context()
	resp, err = ph.productService.GetCustomerProducts(ctx, req)

	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "success",
		Data:    resp,
	})
}
