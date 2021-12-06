package v1

import (
	"profile/common"
	"profile/httpserver/api/v1/middlewares"
	"profile/logic"
	"profile/logic/profile"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "profile/docs"
)

type App struct {
	Logic logic.Logic
}

func New() (*App, error) {
	l := profile.New()
	app := new(App)
	app.Logic = l
	return app, nil
}

func (a *App) ApplyEndpoints(r *chi.Mux) {
	r.Route(common.ServiceConfig.Basepath, func(r chi.Router) {
		r.Use(
			middleware.RequestID,
			common.NewStructuredLogger(common.Logger),
			middleware.Recoverer,
			middlewares.AddContentType,
		)
		r.Mount("/debug", profiler())
		r.Get("/i", info)
		r.Post("/upload_template", uploadTemplate)
		r.Post("/status", receiveStatus)
		r.Route("/receivers", func(r chi.Router) {
			r.Post("/{user_id}", a.createReceiversOnUser)
			r.Get("/{user_id}", a.readReceiversOnUser)
			r.Put("/{user_id}", a.updateReceiversOnUser)
			r.Delete("/{user_id}", a.deleteReceiversOnUser)
			r.Get("/", a.readAllReceivers)
		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
}
