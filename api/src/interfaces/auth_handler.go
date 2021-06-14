package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type AuthHandler struct {
	ua application.UserAppInterface
	au application.AuthAppInterface
}

func NewAuthHandler(ua application.UserAppInterface, au application.AuthAppInterface) *AuthHandler {
	return &AuthHandler{
		ua: ua,
		au: au,
	}
}

func (ah *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var lu *entity.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&lu); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
	}
	defer r.Body.Close()

	user := entity.User{
		Email:    lu.Email,
		Password: lu.Password,
	}

	result, err := ah.ua.LoginUser(user)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	rt := result["refresh_token"].(*entity.RefreshToken)
	//save result["refreshToken"] to redis metadata
	if authErr := ah.au.CreateAuth(rt); authErr != nil {
		w.WriteHeader(authErr.Status)
		w.Write(authErr.ResponseJSON().([]byte))
		return
	}

	//respose payload(user, accessToken)
	pUser := result["user"].(*entity.User)
	at := result["access_token"].(*entity.AccessToken)
	jsonData, jsonErr := json.Marshal(map[string]interface{}{
		"user":          pUser.PublicUser(),
		"access_token":  at,
		"refresh_token": rt,
	})

	if jsonErr != nil {
		jsonErr := errors.NewInternalServerError("internal marshaling error")
		w.WriteHeader(jsonErr.Status)
		w.Write(jsonErr.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}