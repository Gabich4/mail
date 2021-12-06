package middlewares

import (
	"fmt"
	"net/http"
	"profile/common"
	"profile/utils"
)

func CheckEnableDebug(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if enableDebug := r.URL.Query().Get("enable"); enableDebug == "true" {
			next.ServeHTTP(w, r)
		} else {
			err := fmt.Errorf("debug is disabled")
			common.Logger.Print(err)
			w.WriteHeader(http.StatusNotAcceptable)
			utils.SerializeResponseJSON(w, err)
		}
	})
}
