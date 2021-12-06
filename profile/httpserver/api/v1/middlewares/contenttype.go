package middlewares

import "net/http"

func AddContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ct := w.Header().Get("content-type"); ct == "" {
			w.Header().Set("content-type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
