package config

import (
	"eniqlo/package/lumen"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTSecret() string {
	if os.Getenv("JWT_SECRET") == "" {
		return "blindedbypassion"
	}
	return os.Getenv("JWT_SECRET")
}

func JWTConfig() echojwt.Config {
	config := echojwt.Config{
		SigningKey: []byte(JWTSecret()),
		ErrorHandler: func(c echo.Context, err error) error {
			return lumen.FromError(lumen.NewError(lumen.ErrUnauthorized, err)).SendResponse(c)
		},
	}
	return config
}
