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

		router.Route("/auth", func(router chi.Router) {
			router.Post("/login", r.wrapper.Wrap(r.controller.Login))
			router.Post("/refresh", r.wrapper.Wrap(r.controller.Refresh))
			router.Post("/logout", r.wrapper.Wrap(r.controller.Logout))
		})

		router.Route("/databases", func(router chi.Router) {
			router.Get("/", r.wrapper.Wrap(r.controller.GetDatabases)) // Получить все СУБД
			router.Route("/{database_id}", func(router chi.Router) {
				router.Get("/datatypes", r.wrapper.Wrap(r.controller.GetDatatypes)) // Получить все типы данных для СУБД
			})
		})

		router.Route("/diagrams", func(router chi.Router) {
			router.Get("/", r.wrapper.Wrap(r.controller.GetRequestersDiagrams)) // Получить список своих диаграмм
			router.Post("/", r.wrapper.Wrap(r.controller.CreateDiagram))        // Создать диаграмму
			router.Route("/{diagram_id}", func(router chi.Router) {
				router.Get("/", r.wrapper.Wrap(r.controller.GetDiagram))       // Получить полную диаграмму
				router.Put("/", r.wrapper.Wrap(r.controller.EditDiagram))      // Редактировать данные дигараммы
				router.Delete("/", r.wrapper.Wrap(r.controller.DeleteDiagram)) // Удалить диаграмму
				router.Route("/access", func(router chi.Router) {
					router.Get("/", r.wrapper.Wrap(r.controller.GetAccessRights))   // Получить список прав для диаграммы
					router.Post("/", r.wrapper.Wrap(r.controller.GiveAccessRights)) // Выдать права на диаграмму
				})
				router.Route("/tables", func(router chi.Router) {
					router.Post("/", r.wrapper.Wrap(r.controller.CreateTable)) // Создать таблицу в диаграмме
				})
			})
			router.Route("/access", func(router chi.Router) {
				router.Route("/{access_id}", func(router chi.Router) {
					router.Delete("/", r.wrapper.Wrap(r.controller.DeleteAccessRights)) // Удалить права для диаграммы
				})
			})
			router.Route("/tables", func(router chi.Router) {
				router.Route("/{table_id}", func(router chi.Router) {
					router.Put("/", r.wrapper.Wrap(r.controller.EditTable))      // Изменить таблицу
					router.Delete("/", r.wrapper.Wrap(r.controller.DeleteTable)) // Удалить таблицу
					router.Route("/coordinates", func(router chi.Router) {
						router.Get("/", r.wrapper.Wrap(r.controller.GetTableCoordinates))  // Получить координаты таблицы
						router.Put("/", r.wrapper.Wrap(r.controller.EditTableCoordinates)) // Переместить таблицу
					})
					router.Route("/rows", func(router chi.Router) {
						router.Post("/", r.wrapper.Wrap(r.controller.CreateRow)) // Создать строку в таблице
					})
				})
				router.Route("/rows", func(router chi.Router) {
					router.Route("/{row_id}", func(router chi.Router) {
						router.Put("/", r.wrapper.Wrap(r.controller.EditRow))      // Изменить строку
						router.Delete("/", r.wrapper.Wrap(r.controller.DeleteRow)) // Удалить строку
					})
				})
				router.Route("/relations", func(router chi.Router) {
					router.Post("/", r.wrapper.Wrap(r.controller.CreateRelation)) // Создать связь
					router.Route("/{relation_id}", func(router chi.Router) {
						router.Put("/", r.wrapper.Wrap(r.controller.EditRelation))      // Изменить связь
						router.Delete("/", r.wrapper.Wrap(r.controller.DeleteRelation)) // Удалить связь
					})
				})
			})
		})

		router.Route("/users", func(router chi.Router) {
			router.Route("/me", func(router chi.Router) {
				router.Get("/", r.wrapper.Wrap(r.controller.GetRequesterUser))  // Получить данные о себе
				router.Put("/", r.wrapper.Wrap(r.controller.EditRequesterUser)) // Изменить данные о себе
			})
			router.Route("/{user_id}", func(router chi.Router) {
				router.Get("/", r.wrapper.Wrap(r.controller.GetUser)) // Получить пользователя по id
			})
		})

	})
}
