package staffHandler

import (
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/package/lumen"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uh staffHandler) Register(c echo.Context) error {
	var (
		req  request.StaffRegister
		resp *response.UserAccessToken
		err  error
	)
	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	// Validate the User struct
	err = uh.val.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	resp, err = uh.staffService.Register(c.Request().Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusCreated, response.Common{
		Message: "User registered successfully",
		Data:    resp,
	})
}
