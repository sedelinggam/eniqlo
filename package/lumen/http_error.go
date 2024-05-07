package lumen

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func FromError(err error) APIError {
	var apiError APIError
	var svcError Error
	if errors.As(err, &svcError) {
		if svcError.appError == nil {
			apiError.Message = "generic error"
		}
		apiError.Message = svcError.appError.Error()
		svcErr := svcError.SvcError()
		switch svcErr {
		case ErrBadRequest:
			apiError.Status = http.StatusBadRequest
		case ErrUnauthorized:
			apiError.Status = http.StatusUnauthorized
		case ErrNotFound:
			apiError.Status = http.StatusNotFound
		case ErrConflict:
			apiError.Status = http.StatusConflict
		case ErrInternalFailure:
			apiError.Status = http.StatusInternalServerError
		}

	}
	return apiError
}

func (apiErr APIError) SendResponse(c echo.Context) error {
	return c.JSON(apiErr.Status, echo.Map{
		"error": apiErr.Message,
	})
}
