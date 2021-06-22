package helpers

import (
	"fmt"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func SetJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func ExtractUintParam(r *http.Request, param string) (uint64, *errors.RestErr) {
	value, err := strconv.ParseUint(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError(fmt.Sprintf("%s is not valid", param))
	}
	return value, nil
}

func ExtractIntParam(r *http.Request, param string) (int64, *errors.RestErr) {
	value, err := strconv.ParseInt(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError(fmt.Sprintf("%s is not valid", param))
	}
	return value, nil
}

func ExtractStringParam(r *http.Request, param string) string {
	value := chi.URLParam(r, param)
	return value
}
