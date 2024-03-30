package router

import (
	"fmt"
	"github.com/free-diagrams/sql-backend/internal/api/http/controller"
	"github.com/free-diagrams/sql-backend/internal/api/http/middleware"
	"github.com/free-diagrams/sql-backend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	router     *gin.Engine
	controller *controller.Controller
	log        *zerolog.Logger
}

func New(controller *controller.Controller, log *zerolog.Logger) *Router {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.ZerologMiddleware(*log))
	router.Use(middleware.ProduceLanguageMiddleware())
	return &Router{
		router:     router,
		controller: controller,
		log:        log,
	}
}

func (r *Router) RegisterRoutes() {
	api := r.router.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		roots := api.Group("/colors")
		{
			roots.GET("", r.controller.GetColors)
		}
	}
}

func (r *Router) StartHTTPServer(cfg config.HTTPConfig) error {
	err := r.router.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}
