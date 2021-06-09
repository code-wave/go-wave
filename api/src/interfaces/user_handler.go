package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/domain/entity"
)

type UserHandler struct {
	ua application.UserAppInterface
}

func NewUserHandler(ua application.UserAppInterface) *UserHandler {
	return &UserHandler{
		ua: ua,
	}
}

func (uh *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	var u entity.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid json body", http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	newUser, err := uh.ua.SaveUser(&u)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(newUser.ResponseJSON().([]byte))
}
