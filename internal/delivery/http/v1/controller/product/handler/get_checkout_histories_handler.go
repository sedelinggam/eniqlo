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

func (ph productHandler) GetCheckoutHistories(c echo.Context) error {

	var (
		req  request.GetCheckoutHistories
		resp *[]response.GetCheckoutHistories
		err  error
	)

	queries := c.QueryParams()

	shouldFilter := request.ShouldGetCheckoutHistoriesFilter{
		CustomerID: queries.Has("customerId"),
		Limit:      queries.Has("limit"),
		Offset:     queries.Has("offset"),
		CreatedAt:  queries.Has("createdAt"),
	}

	req = request.GetCheckoutHistories{
		CustomerID: c.QueryParam("customerId"),
	}

	if customerID := c.QueryParam("customerID"); customerID != "" {
		err := ph.val.Var(c.QueryParam("customerID"), "uuid")
		if err == nil {
			req.CustomerID = c.QueryParam("customerID")
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

	ctx := context.WithValue(c.Request().Context(), request.ShouldGetCheckoutHistoriesFilterKey, shouldFilter)

	resp, err = ph.checkoutService.GetCheckoutHistories(ctx, req)

	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "success",
		Data:    resp,
	})

}
