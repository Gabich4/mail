package main

import (
	"auth/config"
	"auth/utils"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

var cfg = config.GetConfig()

//// Create a struct to read the username and password from the request body
//type Credentials struct {
//	Password string `json:"password"`
//	Username string `json:"username"`
//}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Use(utils.NewStructuredLogger(utils.Logger))
	r.Use(addContentType)

	r.Mount("/debug", Profiler())

	r.Group(func(r chi.Router) {
		r.Use(redirectMiddleware)
		r.Get("/", login)
		r.Get("/login", login)
		r.Get("/logout", logout)
	})

	r.Group(func(r chi.Router) {
		r.Use(checkAccessAndRefreshCookiesMiddleware)
		r.Get("/i", info)
		r.Get("/me", info)
	})

	cfg := config.GetConfig()
	addr := fmt.Sprintf("localhost:%v", cfg.Port)
	utils.Logger.Printf("Service auth started. %v", addr)

	utils.Logger.Fatal(http.ListenAndServe(addr, r))
}
