package helpers

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func SetJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func ExtractUintParam(r *http.Request, param string) (uint64, error) {
	value, err := strconv.ParseUint(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("%s is not valid", param))
	}
	return value, nil
}

func ExtractIntParam(r *http.Request, param string) (int64, error) {
	value, err := strconv.ParseInt(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("%s is not valid", param))
	}
	return value, nil
}
