package v1

import (
	"bodyshop/common"
	"bodyshop/logic"
	"bodyshop/logic/bodyshop"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "bodyshop/docs"

	"github.com/go-chi/chi/v5"
)

// App is a structure that represents logic layer.
// It contains a client to mailsender service.
type App struct {
	Logic            logic.Logic
	MailsenderClient http.Client
}

func New() *App {
	logic := bodyshop.New()
	msClient := http.DefaultClient

	app := new(App)
	app.Logic = logic
	app.MailsenderClient = *msClient

	return app
}

// ApplyEndpoints sets routes to router.
func (a *App) ApplyEndpoints(r *chi.Mux) {
	r.Route(common.ServiceConfig.Basepath, func(r chi.Router) {
		r.Use(common.NewStructuredLogger(common.Logger))
		r.Post("/send", a.send)
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
}
