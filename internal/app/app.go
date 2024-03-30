package app

import (
	"encoding/json"
	"fmt"
	"github.com/free-diagrams/sql-backend/internal/api/http/controller"
	"github.com/free-diagrams/sql-backend/internal/api/http/router"
	"github.com/free-diagrams/sql-backend/internal/config"
	"github.com/free-diagrams/sql-backend/internal/repository/postgres"
	"github.com/free-diagrams/sql-backend/internal/repository/postgres/adapter/color_repository"
	"github.com/free-diagrams/sql-backend/internal/service/color_service"
	"github.com/free-diagrams/sql-backend/internal/usecase"
	"github.com/free-diagrams/sql-backend/pkg/core"
	"github.com/free-diagrams/sql-backend/pkg/loclzr"
	"github.com/free-diagrams/sql-backend/pkg/logging"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"golang.org/x/text/language"
	"os"
)

func Run() {
	cfg, err := config.Parse()
	if err != nil {
		panic(errors.Wrap(err, "failed to parse config"))
	}

	log, err := createLogger(cfg)
	if err != nil {
		panic(errors.Wrap(err, "failed to create logger"))
	}
	log.Debug().Msg("Logger initialized")

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
	bundle.MustLoadMessageFile("/locale/en-US.json")
	bundle.MustLoadMessageFile("/locale/ru-RU.json")
	localizer := loclzr.New(bundle)

	colorRepository := color_repository.New(db, log)

	colorService := color_service.New(colorRepository, localizer, log)

	useCase := usecase.New(colorService, log)

	controller := controller.New(useCase, localizer, log)

	httpRouter := router.New(controller, log)
	httpRouter.RegisterRoutes()
	err = httpRouter.StartHTTPServer(cfg.HTTP)
	if err != nil {
		panic(errors.Wrap(err, "failed to start http server"))
	}
}

func createLogger(cfg *config.Config) (*zerolog.Logger, error) {
	level := logging.ParseZerologLevel(cfg.Logger.Level)

	switch cfg.App.Environment {
	case core.EnvProd, core.EnvTest:
		logger := zerolog.New(os.Stdout).Level(level).With().
			Str("service", cfg.App.Name).Timestamp().Caller().Logger()
		return &logger, nil
	case core.EnvDev:
		consoleWriter := zerolog.NewConsoleWriter()
		logger := zerolog.New(consoleWriter).Level(level).With().
			Str("service", cfg.App.Name).Timestamp().Caller().Logger()
		return &logger, nil
	default:
		return nil, fmt.Errorf("failed to initialize logger")
	}
}
