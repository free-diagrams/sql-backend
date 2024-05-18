package app

import (
	"encoding/json"
	routerdoc "github.com/free-diagrams/sql-backend/internal/api/http/doc/router"
	"github.com/free-diagrams/sql-backend/internal/api/http/middleware"
	"github.com/free-diagrams/sql-backend/internal/api/http/responsewriter"
	"github.com/free-diagrams/sql-backend/internal/api/http/v1/controller"
	routerv1 "github.com/free-diagrams/sql-backend/internal/api/http/v1/router"
	"github.com/free-diagrams/sql-backend/internal/api/http/wrapper"
	"github.com/free-diagrams/sql-backend/internal/config"
	"github.com/free-diagrams/sql-backend/internal/infrastructure/repository/postgres"
	"github.com/free-diagrams/sql-backend/internal/infrastructure/repository/postgres/adapter"
	"github.com/free-diagrams/sql-backend/internal/infrastructure/repository/postgres/txmanager"
	usecasev1 "github.com/free-diagrams/sql-backend/internal/usecase/v1"
	"github.com/free-diagrams/sql-backend/pkg/chirouter"
	"github.com/free-diagrams/sql-backend/pkg/httpserver"
	"github.com/free-diagrams/sql-backend/pkg/loclzr"
	"github.com/free-diagrams/sql-backend/pkg/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
	"golang.org/x/text/language"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg, err := config.Parse()
	if err != nil {
		panic(errors.Wrap(err, "failed to parse config"))
	}

	log, err := logger.NewConsoleAndHookLogger(cfg.Logger.Level)
	if err != nil {
		panic(errors.Wrap(err, "failed to create logger"))
	}
	log.Debug().Msg("Logger initialized")
	goose.SetLogger(log)

	db, err := postgres.New(cfg.DB)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect database"))
	}
	log.Debug().Msg("Database connected")
	defer func() {
		err = db.Close()
		if err != nil {
			log.Warn().Msg("Failed to close database connection")
		}
	}()

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	bundle.MustLoadMessageFile("internal/domain/locale/en-US.json")
	bundle.MustLoadMessageFile("internal/domain/locale/ru-RU.json")
	localizer := loclzr.New(bundle)

	txManager := txmanager.New(db, log)

	repository := adapter.New(db, log)

	useCase := usecasev1.New(txManager, repository, repository, log)

	responseWriter := responsewriter.New(localizer, log)

	controller := controller.New(useCase, responseWriter, log)

	wrapper := wrapper.New(responseWriter, log)

	middleware := middleware.New(useCase, responseWriter, log)

	chiMux := chirouter.NewChiMux()

	docRouter := routerdoc.New(chiMux, log)
	v1Router := routerv1.New(chiMux, middleware, wrapper, controller, log)

	docRouter.RegisterRoutes()
	v1Router.RegisterRoutes()

	httpServer := httpserver.New(
		chiMux,
		httpserver.Addr(cfg.HTTP.Host, cfg.HTTP.Port),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Str("signal", s.String()).Msg("signal received")
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Error().Err(err).Msg("failed to shutdown server")
	}
}
