package application

import (
	"log"

	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/auth"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type AuthApp struct {
	ar repository.AuthRepository
}

type AuthAppInterface interface {
	CreateAuth(*entity.RefreshToken) *errors.RestErr
	DeleteAuth(string) *errors.RestErr
	FetchAuth(string) (int64, *errors.RestErr)
	Refresh(string, int64) (*entity.AccessToken, *errors.RestErr)
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

func (au *AuthApp) FetchAuth(uuid string) (int64, *errors.RestErr) {
	return au.ar.Fetch(uuid)
}

func (au *AuthApp) Refresh(uuid string, uid int64) (*entity.AccessToken, *errors.RestErr) {
	userID, err := au.ar.Fetch(uuid)
	if err != nil {
		return nil, err
	}

	if userID != uid {
		authErr := errors.NewUnauthorizedError("unauthroized, userID which is parsed from access_token is diffrent redis's userID")
		log.Printf("access_token's userID: %d but redis's userID: %d invalid user access", uid, userID)
		return nil, authErr
	}
	log.Println(userID, uid)

	at, tokenErr := auth.JwtWrapper.GenerateAccessToken(userID)
	if tokenErr != nil {
		restErr := errors.NewInternalServerError("token generation error")
		return nil, restErr
	}

	return at, nil
}
