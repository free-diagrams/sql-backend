package middleware

import (
	"github.com/free-diagrams/sql-backend/internal/api/http/responsewriter"
	"github.com/free-diagrams/sql-backend/internal/domain/usecase"
	"github.com/free-diagrams/sql-backend/pkg/logger"
)

type Middleware struct {
	userUseCase    usecase.UserUseCase
	responseWriter *responsewriter.ResponseWriter
	log            *logger.Logger
}

func New(userUseCase usecase.UserUseCase, responseWriter *responsewriter.ResponseWriter, log *logger.Logger) *Middleware {
	return &Middleware{
		userUseCase:    userUseCase,
		responseWriter: responseWriter,
		log:            log,
	}
}
