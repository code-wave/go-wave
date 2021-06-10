package repository

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type UserRepository interface {
	Save(*entity.User) *errors.RestErr
	Get(*entity.User) *errors.RestErr
	GetAll(int64, int64) (entity.Users, *errors.RestErr)
	Update(*entity.User) *errors.RestErr
	Delete(uint64) *errors.RestErr
	FindByEmailAndPassword(*entity.User) (*entity.User, *errors.RestErr)
}
