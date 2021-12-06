package service

import (
	"bodyshop/common"
	v1 "bodyshop/httpserver/api/v1"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Service is a structure representing service.
// It contains HTTP-server and App layer.
type Service struct {
	Server *http.Server
	App    *v1.App
}

func New() *Service {
	app := v1.New()

	r := chi.NewRouter()
	app.ApplyEndpoints(r)

	// TODO: fetch info from config
	server := http.Server{
		Addr:    common.ServiceConfig.Host,
		Handler: r,
	}

	service := new(Service)
	service.App = app
	service.Server = &server

	return service
}
