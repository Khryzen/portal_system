package api

import (
	"net/http"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	api := strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(api, "confirm_login") {
		ConfirmLogin(w, r)
	}
}
