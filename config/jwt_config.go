package config

import (
	"os"
)

func JWTSecret() string {
	if os.Getenv("JWT_SECRET") == "" {
		return "blindedbypassion"
	}
	return os.Getenv("JWT_SECRET")
}
