package router

import (
	"github.com/free-diagrams/sql-backend/pkg/logger"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	router *chi.Mux
	log    *logger.Logger
}

func New(router *chi.Mux, log *logger.Logger) *Router {
	return &Router{
		router: router,
		log:    log,
	}
}

func (r *Router) RegisterRoutes() {
	r.router.Route("/api/docs", func(router chi.Router) {
		router.Mount("/swagger", httpSwagger.WrapHandler)
	})
}
