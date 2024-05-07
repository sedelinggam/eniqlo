package cache

import "eniqlo/internal/delivery/http/v1/response"

func AddAccessToken(userID string, sv *response.UserAccessToken) {
	if accessToken == nil {
		caT := make(cacheAccessToken)
		accessToken = &caT
	}

	(*accessToken)[userID] = AccessToken{
		JWTClaim: *sv,
	}
}

func GetShortVideo(userID string) *AccessToken {
	if accessToken == nil {
		return nil
	}

	if val, ok := (*accessToken)[userID]; ok {
		return &val
	}
	return nil
}
