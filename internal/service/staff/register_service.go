package staffService

import (
	"context"
	"eniqlo/internal/delivery/http/v1/request"
	"eniqlo/internal/delivery/http/v1/response"
	"eniqlo/internal/entity"
	cache "eniqlo/package/caching"
	"eniqlo/package/crypto/bcrypt"
	cryptoJWT "eniqlo/package/crypto/jwt"
	"eniqlo/package/lumen"
	"errors"
	"time"

	"github.com/google/uuid"
)

func (ss staffService) Register(ctx context.Context, requestData request.StaffRegister) (*response.UserAccessToken, error) {
	var (
		err          error
		hashPassword string
	)

	//Check if user already exist
	cacheAccessToken := cache.GetAccessToken(requestData.PhoneNumber)
	if cacheAccessToken != nil && cacheAccessToken.Expired.After(time.Now()) {
		return nil, lumen.NewError(lumen.ErrConflict, errors.New("conflicts"))
	}

	//Password Hash
	hashPassword, err = bcrypt.HashPassword(requestData.Password)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	//Create User
	userData := entity.Staff{
		ID:          uuid.New().String(),
		PhoneNumber: requestData.PhoneNumber,
		Name:        requestData.Name,
		Password:    hashPassword,
		CreatedAt:   time.Now(),
	}

	//Check Phone Number
	err = userData.CheckPhoneNumber()
	if err != nil {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	err = ss.staffRepo.Create(ctx, userData)
	if err != nil {
		//Duplicate unique key
		if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	// Create the Claims
	accessToken, err := cryptoJWT.GenerateToken(userData.ID, userData.PhoneNumber)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Cache The token
	respAccessToken := &response.UserAccessToken{
		PhoneNumber: requestData.PhoneNumber,
		Name:        requestData.Name,
		AccessToken: *accessToken,
	}

	cache.AddAccessToken(userData.PhoneNumber, respAccessToken, time.Now().Add(time.Hour*7))

	return respAccessToken, nil
}
