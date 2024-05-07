package cryptoJWT

import (
	"eniqlo/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Id          string
	PhoneNumber string
	jwt.RegisteredClaims
}

func GenerateToken(id, phoneNumber string) (*string, error) {
	secret := config.JWTSecret()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Id:          id,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
	})

	tokenString, err := token.SignedString([]byte(secret))
	return &tokenString, err
}

type JWTPayload struct {
	Id          string
	PhoneNumber string
}

func VerifyToken(token string) (*JWTPayload, error) {
	secret := config.JWTSecret()

	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims.RegisteredClaims.ExpiresAt.Before(time.Now()) {
		return nil, err
	}

	payload := &JWTPayload{
		Id:          claims.Id,
		PhoneNumber: claims.PhoneNumber,
	}

	return payload, nil
}
