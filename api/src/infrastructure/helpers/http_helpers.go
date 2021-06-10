package helpers

import "net/http"

func JsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
