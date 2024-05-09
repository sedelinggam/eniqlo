package productHandler

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/package/lumen"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ph productHandler) GetProducts(c echo.Context) error {
	var (
		req  request.GetProducts
		resp *[]response.GetProducts
		err  error
	)

	queries := c.QueryParams()
	shouldFilter := request.ShouldGetProductsFilter{
		ID:          queries.Has("id"),
		Name:        queries.Has("name"),
		Sku:         queries.Has("sku"),
		Category:    queries.Has("category"),
		Price:       queries.Has("price"),
		IsAvailable: queries.Has("isAvailable"),
		CreatedAt:   queries.Has("createdAt"),
		InStock:     queries.Has("inStock"),
	}

	req = request.GetProducts{
		Name: c.QueryParam("name"),
		Sku:  c.QueryParam("sku"),
	}

	if name := c.QueryParam("name"); name != "" {
		req.Name = c.QueryParam("name")
	}

	if id := c.QueryParam("id"); id != "" {
		err := ph.val.Var(c.QueryParam("id"), "uuid")
		if err == nil {
			req.ID = c.QueryParam("id")
		}
	}

	if isAvailable := c.QueryParam("isAvailable"); isAvailable != "" {
		err := ph.val.Var(c.QueryParam("isAvailable"), "boolean")
		if err == nil {
			val, _ := strconv.ParseBool(c.QueryParam("isAvailable"))
			req.IsAvailable = val
		}
	}

	if category := c.QueryParam("category"); category != "" {
		err := ph.val.Var(c.QueryParam("category"), "oneof=Clothing Accessories Footwear Beverages")
		if err == nil {
			req.Category = c.QueryParam("category")
		}
	}

	if price := c.QueryParam("price"); price != "" {
		err := ph.val.Var(c.QueryParam("price"), "oneof=asc desc")
		if err == nil {
			req.Price = c.QueryParam("price")
		}
	}

	if inStock := c.QueryParam("inStock"); inStock != "" {
		err := ph.val.Var(c.QueryParam("inStock"), "boolean")
		if err == nil {
			val, _ := strconv.ParseBool(c.QueryParam("inStock"))
			req.InStock = val
		}
	}

	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		err := ph.val.Var(c.QueryParam("createdAt"), "oneof=asc desc")
		if err == nil {
			req.CreatedAt = c.QueryParam("createdAt")
		}
	}

	if limit := c.QueryParam("limit"); limit != "" {
		err := ph.val.Var(c.QueryParam("limit"), "number")
		if err == nil {
			val, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 32)
			req.Limit = int32(val)
		}
	} else {
		req.Limit = 5
	}

	if offset := c.QueryParam("offset"); offset != "" {
		err := ph.val.Var(c.QueryParam("offset"), "number")
		if err == nil {
			val, _ := strconv.ParseInt(c.QueryParam("offset"), 10, 32)
			req.Offset = int32(val)
		}
	} else {
		req.Offset = 0
	}

	ctx := context.WithValue(c.Request().Context(), request.ShouldGetProductsFilterKey, shouldFilter)

	resp, err = ph.productService.GetProducts(ctx, req)

	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "success",
		Data:    resp,
	})
}
