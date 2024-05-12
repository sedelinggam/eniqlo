package productHandler

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	cryptoJWT "eniqlo/package/crypto/jwt"
	"eniqlo/package/lumen"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (ph productHandler) CheckoutProduct(c echo.Context) error {
	var (
		req  request.CheckoutProduct
		resp *response.CheckoutResponse
		err  error
	)

	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	//Get jwt user ID
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cryptoJWT.JWTClaims)
	phoneNumber := claims.PhoneNumber
	ctx := context.WithValue(c.Request().Context(), cryptoJWT.KeyPhoneNumber, phoneNumber)

	// Create a new validator instance
	err = ph.val.Struct(req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}
	if req.Change == nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, errors.New("change nill"))).SendResponse(c)
	}
	for _, v := range req.ProductDetails {
		err = ph.val.Struct(v)
		if err != nil {
			return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
		}
	}

	resp, err = ph.checkoutService.CheckoutProduct(ctx, req)

	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "success",
		Data:    resp,
	})
}
