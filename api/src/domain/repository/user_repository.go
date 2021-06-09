package repository

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type UserRepository interface {
	Save(*entity.User) (*entity.User, *errors.RestErr)
}
