package productHandler

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	valueobject "eniqlo/internal/value_object"
	cryptoJWT "eniqlo/package/crypto/jwt"
	"eniqlo/package/lumen"
	"errors"
	"net/http"
	"regexp"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (ph productHandler) UpdateProduct(c echo.Context) error {
	var (
		req  request.UpdateProduct
		resp *response.UpdateProduct
		err  error
	)
	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	if id := c.Param("id"); id != "" {
		err := ph.val.Var(c.Param("id"), "uuid")
		if err != nil {
			return lumen.FromError(lumen.NewError(lumen.ErrNotFound, err)).SendResponse(c)
		}
		req.ID = c.Param("id")
	}

	// Create a new validator instance
	err = ph.val.Struct(req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	//Check product category
	err = valueobject.CheckProductCategory(req.Category)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	urlRegex := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	var re = regexp.MustCompile(urlRegex)
	if !re.MatchString(req.ImageURL) {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, errors.New("invalid url"))).SendResponse(c)
	}

	//Get jwt user ID
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cryptoJWT.JWTClaims)
	phoneNumber := claims.PhoneNumber
	ctx := context.WithValue(c.Request().Context(), cryptoJWT.KeyPhoneNumber, phoneNumber)

	resp, err = ph.productService.UpdateProduct(ctx, req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "success",
		Data:    resp,
	})
}
