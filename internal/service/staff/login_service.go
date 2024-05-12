package staffService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/package/crypto/bcrypt"
	cryptoJWT "eniqlo/package/crypto/jwt"
	"eniqlo/package/lumen"
	"errors"
)

func (ss staffService) Login(ctx context.Context, requestData request.StaffLogin) (*response.UserAccessToken, error) {
	//Password Hash
	var (
		err error
	)

	// Find the user by credentials
	user, err := ss.staffRepo.GetStaffByPhoneNumber(ctx, requestData.PhoneNumber)
	if err != nil {
		//Duplicate unique key
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Compare password hash
	if !bcrypt.CheckPasswordHash(requestData.Password, user.Password) {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("password doesn't match"))
	}
	// Create the Claims
	accessToken, err := cryptoJWT.GenerateToken(user.ID, user.PhoneNumber)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	respAccessToken := &response.UserAccessToken{
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
		AccessToken: *accessToken,
	}

	return respAccessToken, nil
}
