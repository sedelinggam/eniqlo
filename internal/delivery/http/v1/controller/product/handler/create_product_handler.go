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

// @Summary Create Product
// @Description create product
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CreateProduct
// @Param request.CreateProduct
// @Header 200 {string} Token "Token"
// @Router /v1/product [post]
func (ph productHandler) CreateProduct(c echo.Context) error {
	var (
		req  request.CreateProduct
		resp *response.CreateProduct
		err  error
	)
	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
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

	//Get jwt user ID
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cryptoJWT.JWTClaims)
	phoneNumber := claims.PhoneNumber
	ctx := context.WithValue(c.Request().Context(), cryptoJWT.KeyPhoneNumber, phoneNumber)

	//Regex
	urlRegex := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	var re = regexp.MustCompile(urlRegex)
	if !re.MatchString(req.ImageURL) {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, errors.New("invalid url"))).SendResponse(c)
	}

	resp, err = ph.productService.CreateProduct(ctx, req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusCreated, response.Common{
		Message: "success",
		Data:    resp,
	})
}
