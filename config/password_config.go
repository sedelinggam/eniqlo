package config

import (
	"os"
	"strconv"
)

func PasswordCost() int {
	if os.Getenv("BCRYPT_SALT") == "" {
		return 10
	}
	salt, _ := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	return salt
}
