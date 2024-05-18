package router

import (
	"github.com/free-diagrams/sql-backend/internal/api/http/middleware"
	"github.com/free-diagrams/sql-backend/internal/api/http/v1/controller"
	"github.com/free-diagrams/sql-backend/internal/api/http/wrapper"
	"github.com/free-diagrams/sql-backend/pkg/logger"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	router     *chi.Mux
	middleware *middleware.Middleware
	wrapper    *wrapper.Wrapper
	controller *controller.Controller
	log        *logger.Logger
}

func New(router *chi.Mux, middleware *middleware.Middleware, wrapper *wrapper.Wrapper,
	controller *controller.Controller, log *logger.Logger) *Router {
	return &Router{
		router:     router,
		middleware: middleware,
		wrapper:    wrapper,
		controller: controller,
		log:        log,
	}
}

func (r *Router) RegisterRoutes() {
	r.router.Route("/v1", func(router chi.Router) {

		router.Use(r.middleware.AcceptLanguageMiddleware)
		router.Use(r.middleware.AuthMiddleware)

		router.Route("/databases", func(router chi.Router) {
			router.Get("/", r.wrapper.Wrap(r.controller.GetDatabases))
		})

	})
}
