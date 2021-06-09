package application

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

var _ UserAppInterface = &UserApp{}

type UserApp struct {
	ur repository.UserRepository
}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, *errors.RestErr)
}

func NewUserApp(ur repository.UserRepository) *UserApp {
	return &UserApp{
		ur: ur,
	}
}

func (ua *UserApp) SaveUser(user *entity.User) (*entity.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	//encrypt password and set created_at(string)
	if err := user.BeforeSave(); err != nil {
		return nil, err
	}

	return ua.ur.Save(user)
}
