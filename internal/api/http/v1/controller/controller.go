package controller

import (
	"github.com/free-diagrams/sql-backend/internal/api/http/responsewriter"
	"github.com/free-diagrams/sql-backend/internal/domain/usecase"
	"github.com/free-diagrams/sql-backend/pkg/logger"
)

type Controller struct {
	databaseUseCase usecase.DatabaseUseCase
	responseWriter  *responsewriter.ResponseWriter
	log             *logger.Logger
}

func New(databaseUseCase usecase.DatabaseUseCase, responseWriter *responsewriter.ResponseWriter, log *logger.Logger) *Controller {
	return &Controller{
		databaseUseCase: databaseUseCase,
		responseWriter:  responseWriter,
		log:             log,
	}
}
