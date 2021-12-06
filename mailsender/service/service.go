package service

import (
	"mailsender/common"
	v1 "mailsender/httpserver/api/v1"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Service is a structure implementating service's server.
type Service struct {
	Server http.Server
	App    *v1.App
}

// New reads data from service config,
// applying API's endpoints to router
// and returns the new Service or error if any
func New() (*Service, error) {
	app, err := v1.New()
	if err != nil {
		return nil, err
	}

	// creating and applying endpoints
	// to router
	router := chi.NewRouter()
	app.ApplyEndpoints(router)

	server := http.Server{
		Addr:    common.ServiceConfig.Host,
		Handler: router,
	}

	return &Service{
		App:    app,
		Server: server,
	}, nil
}
