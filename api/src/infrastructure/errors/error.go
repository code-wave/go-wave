package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string) error {
	return errors.New(message)
}

//400
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_reqeust",
	}
}

//404
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

//500
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

//403
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusForbidden,
		Error:   "forbidden_request",
	}
}

//401
func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "unauthorized",
	}
}

func (e *RestErr) ResponseJSON() interface{} {
	eJSON, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return eJSON
}
