package service

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"profile/common"
	v1 "profile/httpserver/api/v1"
	"time"

	"github.com/go-chi/chi"
)

// Service represents the whole service.
// It contains simple http.Server
// and application layer.
type Service struct {
	Server *http.Server
	App    *v1.App
}

func New() (*Service, error) {
	app, err := v1.New()
	if err != nil {
		return nil, err
	}
	router := chi.NewRouter()
	app.ApplyEndpoints(router)

	server := http.Server{
		Addr:    common.ServiceConfig.Host,
		Handler: router,
	}

	service := new(Service)
	service.App = app
	service.Server = &server

	return service, nil
}

func (s *Service) Serve() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			s.Shutdown()
			break
		}
	}()

	return s.Server.ListenAndServe()
}

func (s *Service) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	common.Logger.Info("shutting down...")
	_ = s.Server.Shutdown(ctx)
}
