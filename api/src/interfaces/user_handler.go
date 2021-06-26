package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

type UserHandler struct {
	ua application.UserAppInterface
}

func NewUserHandler(ua application.UserAppInterface) *UserHandler {
	return &UserHandler{
		ua: ua,
	}
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	userID, err := helpers.ExtractUintParam(r, "user_id")

	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	user, err := uh.ua.GetUser(userID)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(user.ResponseJSON().([]byte))
}

func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	limit, err := helpers.ExtractIntParam(r, "limit")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	offset, err := helpers.ExtractIntParam(r, "offset")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	users, err := uh.ua.GetAllUsers(limit, offset)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", users.ResponseJSON())
	// w.Write(user.ResponseJSON().([]byte))
}

func (uh *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	var u entity.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}
	defer r.Body.Close()

	newUser, err := uh.ua.SaveUser(u)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(newUser.ResponseJSON().([]byte))
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)
	userID, err := helpers.ExtractUintParam(r, "user_id")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	var u entity.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}
	defer r.Body.Close()

	u.ID = userID
	updateUser, err := uh.ua.UpdateUser(u)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(updateUser.ResponseJSON().([]byte))
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	userID, err := helpers.ExtractUintParam(r, "user_id")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}
	if err := uh.ua.DeleteUser(userID); err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	// result, _ := json.Marshal(map[string]string{"result": "success"})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\"result\": \"success\""))
}
