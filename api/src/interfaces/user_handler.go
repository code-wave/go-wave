package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	ua application.UserAppInterface
}

func NewUserHandler(ua application.UserAppInterface) *UserHandler {
	return &UserHandler{
		ua: ua,
	}
}

func getUserID(r *http.Request, param string) (uint64, *errors.RestErr) {
	userID, err := strconv.ParseUint(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user_id param is invalid")
	}
	return userID, nil
}

func getQueryParam(r *http.Request, param string) (int64, *errors.RestErr) {
	value, err := strconv.ParseInt(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("invalid query param")
	}
	return value, nil
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := getUserID(r, "user_id")
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
	w.Header().Set("Content-Type", "application/json")
	limit, err := getQueryParam(r, "limit")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}
	offset, err := getQueryParam(r, "offset")
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
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	userID, err := getUserID(r, "user_id")
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
	w.Header().Set("Content-Type", "application/json")
	userID, err := getUserID(r, "user_id")
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

	result, _ := json.Marshal(map[string]string{"result": "success"})
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
