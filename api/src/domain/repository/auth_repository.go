package repository

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type AuthRepository interface {
	Create(rt *entity.RefreshToken) *errors.RestErr
	Delete()
	Fetch(uid string) (string, *errors.RestErr)
}
