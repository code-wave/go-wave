package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

const ErrNoRows = "no_rows_error"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string) error {
	return errors.New(message)
}

//200 duplicated error
func NewDuplicatedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusOK,
		Error:   "duplicated",
	}
}

//200 wrong information error
func NewWrongInfoError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusOK,
		Error:   "wrong_info",
	}
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

func NewNoRowsError() *RestErr {
	return &RestErr{
		Message: ErrNoRows,
		Status:  http.StatusOK,
		Error:   ErrNoRows,
	}
}

func (e *RestErr) ResponseJSON() interface{} {
	eJSON, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return eJSON
}
