package repository

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type AuthRepository interface {
	Create(*entity.RefreshToken) *errors.RestErr
	Delete(string) *errors.RestErr
	Fetch(string) (string, *errors.RestErr)
}
