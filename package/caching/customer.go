package cache

import (
	"eniqlo/internal/delivery/http/v1/response"
	"time"
)

func AddCustomer(phoneNumber string, sv *response.UserAccessToken, jwtExpired time.Time) {
	if accessToken == nil {
		caT := make(cacheAccessToken)
		accessToken = &caT
	}

	(*accessToken)[phoneNumber] = AccessToken{
		JWTClaim: *sv,
		Expired:  jwtExpired,
	}
}

func GetCustomer(phoneNumber string) *AccessToken {
	if accessToken == nil {
		return nil
	}

	if val, ok := (*accessToken)[phoneNumber]; ok {
		return &val
	}
	return nil
}
