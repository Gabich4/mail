package v1

import (
	"context"
	"mailsender/common"
	"mailsender/logic"
	"mailsender/logic/mailsender"
	"mailsender/repository/mongorepository"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "mailsender/docs"

	"github.com/go-chi/chi/v5"
)

// App structure represents the API layer
type App struct {
	Logic         logic.Logic
	ProfileClient *http.Client
}

// New returns the new instance of App or error if any.
func New() (*App, error) {
	ctx := context.Background()
	repo, err := mongorepository.New(ctx)
	if err != nil {
		return nil, err
	}

	ms := mailsender.New(repo)
	if err != nil {
		return nil, err
	}

	app := new(App)
	app.Logic = ms
	app.ProfileClient = http.DefaultClient

	return app, nil
}

// ApplyEndpoints configures the passed router applying
// http handlers and middlewares to it.
func (a *App) ApplyEndpoints(r *chi.Mux) {
	r.Route(common.ServiceConfig.Basepath, func(r chi.Router) {
		r.Use(common.NewStructuredLogger(common.Logger))
		r.Post("/send", a.send)
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
}
