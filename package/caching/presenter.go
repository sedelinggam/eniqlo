package cache

import "eniqlo/internal/delivery/http/v1/response"

type AccessToken struct {
	JWTClaim response.UserAccessToken
}

type (
	cacheAccessToken map[string]AccessToken
)

var (
	accessToken *cacheAccessToken
)
