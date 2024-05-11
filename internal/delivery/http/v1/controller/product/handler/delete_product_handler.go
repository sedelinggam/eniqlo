package productHandler

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

func (ph productHandler) DeleteProduct(c echo.Context) error {
	var (
		req  request.DeleteProduct
		resp *response.DeleteProduct
		err  error
	)

	if id := c.Param("id"); id != "" {
		err := ph.val.Var(c.Param("id"), "uuid")
		if err != nil {
			return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
		}
		req.ID = c.Param("id")
	}

	// Get jwt user ID
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cryptoJWT.JWTClaims)
	phoneNumber := claims.PhoneNumber
	ctx := context.WithValue(c.Request().Context(), cryptoJWT.KeyPhoneNumber, phoneNumber)

	resp, err = ph.productService.DeleteProducts(ctx, req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "success",
		Data:    resp,
	})
}
