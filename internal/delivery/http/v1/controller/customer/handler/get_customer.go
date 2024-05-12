package customerHandler

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	cryptoJWT "eniqlo/package/crypto/jwt"
	"eniqlo/package/lumen"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (ch customerHandler) GetCustomer(c echo.Context) error {
	var (
		req  request.Customer
		resp []*response.Customer
		err  error
	)
	queries := c.QueryParams()
	if name := queries.Get("name"); name != "" {
		name := queries.Get("name")
		req.Name = &name
	}

	if phoneNumber := queries.Get("phoneNumber"); phoneNumber != "" {
		phoneNumber := queries.Get("phoneNumber")
		phoneNumber = "+" + phoneNumber
		req.PhoneNumber = &phoneNumber
	}

	//Get jwt user ID
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cryptoJWT.JWTClaims)
	phoneNumber := claims.PhoneNumber
	ctx := context.WithValue(c.Request().Context(), cryptoJWT.KeyPhoneNumber, phoneNumber)

	resp, err = ch.customerService.GetCustomers(ctx, req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "success",
		Data:    resp,
	})
}
