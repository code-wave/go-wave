package repository

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type AuthRepository interface {
	FetchValidCode(*entity.ValidEmail) *errors.RestErr
	CreateValidCode(*entity.ValidEmail) *errors.RestErr
	Create(*entity.RefreshToken) *errors.RestErr
	Delete(string) *errors.RestErr
	Fetch(string) (int64, *errors.RestErr)
}
