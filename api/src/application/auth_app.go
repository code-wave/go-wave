package application

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type AuthApp struct {
	ar repository.AuthRepository
}

type AuthAppInterface interface {
	CreateAuth(*entity.RefreshToken) *errors.RestErr
	DeleteAuth(string) *errors.RestErr
	FetchAuth(string) (string, *errors.RestErr)
}

func NewAuthApp(ar repository.AuthRepository) *AuthApp {
	return &AuthApp{
		ar: ar,
	}
}

func (au *AuthApp) CreateAuth(rt *entity.RefreshToken) *errors.RestErr {
	return au.ar.Create(rt)
}
func (au *AuthApp) DeleteAuth(uuid string) *errors.RestErr {
	return au.ar.Delete(uuid)
}
func (au *AuthApp) FetchAuth(uuid string) (string, *errors.RestErr) {
	return au.ar.Fetch(uuid)
}
