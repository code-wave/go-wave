package interfaces

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/code-wave/go-wave/infrastructure/helpers"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/interfaces/middleware"
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
	helpers.SetJsonHeader(w)
	var lu *entity.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&lu); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
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

	//respose payload(user, accessToken, refreshToken)
	pUser := result["user"].(*entity.User)
	at := result["access_token"].(*entity.AccessToken)
	rt := result["refresh_token"].(*entity.RefreshToken)

	atCookie := http.Cookie{
		Name:     "access_token",
		Value:    at.AccessToken,
		HttpOnly: true,
	}

	rtCookie := http.Cookie{
		Name:     "refresh_uuid",
		Value:    rt.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &atCookie)
	http.SetCookie(w, &rtCookie)

	//save result["refreshToken"] to redis metadata
	if authErr := ah.au.CreateAuth(rt); authErr != nil {
		w.WriteHeader(authErr.Status)
		w.Write(authErr.ResponseJSON().([]byte))
		return
	}

	jsonData, jsonErr := json.Marshal(map[string]interface{}{
		"user":         pUser.PublicUser(),
		"access_token": at,
		// "refresh_token": rt,
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

func (ah *AuthHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)
	userID := r.Context().Value(middleware.ContextKeyTokenUserID)
	refreshUuid, err := r.Cookie("refresh_uuid")
	if err != nil {
		restErr := errors.NewBadRequestError("cannot get refresh_uuid from cookie")
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}

	if authErr := ah.au.DeleteAuth(refreshUuid.Value); authErr != nil {
		w.WriteHeader(authErr.Status)
		w.Write(authErr.ResponseJSON().([]byte))
		return
	}

	result, _ := json.Marshal(map[string]string{"result": "success"})
	w.WriteHeader(http.StatusOK)
	w.Write(result)

	log.Println(userID)
}

func (ah *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)
	userID := r.Context().Value(middleware.ContextKeyTokenUserID)
	refreshUuid, err := r.Cookie("refresh_uuid")
	if err != nil {
		restErr := errors.NewBadRequestError("cannot get refresh_uuid from cookie")
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}

	at, authErr := ah.au.Refresh(refreshUuid.Value, userID.(int64))
	if authErr != nil {
		w.WriteHeader(authErr.Status)
		w.Write(authErr.ResponseJSON().([]byte))
		return
	}

	jsonData, jsonErr := json.Marshal(map[string]interface{}{
		"access_token": at,
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
