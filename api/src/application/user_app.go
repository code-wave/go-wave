package application

import (
	"time"

	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/auth"
	"github.com/code-wave/go-wave/infrastructure/encryption"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

var _ UserAppInterface = &UserApp{}

type UserApp struct {
	ur repository.UserRepository
}

type UserAppInterface interface {
	SaveUser(entity.User) (*entity.User, *errors.RestErr)
	GetUser(uint64) (*entity.User, *errors.RestErr)
	GetAllUsers(int64, int64) (entity.Users, *errors.RestErr)
	UpdateUser(entity.User) (*entity.User, *errors.RestErr)
	DeleteUser(uint64) *errors.RestErr
	LoginUser(entity.User) (map[string]interface{}, *errors.RestErr)
}

func NewUserApp(ur repository.UserRepository) *UserApp {
	return &UserApp{
		ur: ur,
	}
}

func (ua *UserApp) SaveUser(user entity.User) (*entity.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	//encrypt password and set created_at(string)
	if err := user.BeforeSave(); err != nil {
		return nil, err
	}

	if err := ua.ur.Save(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (ua *UserApp) GetUser(userID uint64) (*entity.User, *errors.RestErr) {
	user := &entity.User{
		ID: userID,
	}

	if err := ua.ur.Get(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (ua *UserApp) GetAllUsers(limit, offset int64) (entity.Users, *errors.RestErr) {
	return ua.ur.GetAll(limit, offset)
}

func (ua *UserApp) UpdateUser(user entity.User) (*entity.User, *errors.RestErr) {
	user.UpdatedAt.Valid = true
	user.UpdatedAt.String = helpers.GetDateString(time.Now())
	user.Password, _ = encryption.Hash(user.Password)

	if err := ua.ur.Update(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ua *UserApp) DeleteUser(userID uint64) *errors.RestErr {
	return ua.ur.Delete(userID)
}

func (ua *UserApp) LoginUser(lu entity.User) (map[string]interface{}, *errors.RestErr) {
	user, err := ua.ur.FindByEmailAndPassword(&lu)
	if err != nil {
		return nil, err
	}

	token, tokenErr := auth.JwtWrapper.GenerateTokenPair(user.ID)
	if tokenErr != nil {
		restErr := errors.NewInternalServerError("token generation error")
		return nil, restErr
	}

	return map[string]interface{}{
		"user":          user,
		"access_token":  token["access_token"],
		"refresh_token": token["refresh_token"],
	}, nil
}
