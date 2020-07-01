package utils

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetPathVar(key string, r *http.Request) string {
	param := mux.Vars(r)
	return param[key]
}
