package usecase

import (
	"github.com/free-diagrams/sql-backend/internal/domain/service"
	"github.com/rs/zerolog"
)

type UseCase struct {
	colorService service.Color
	log          *zerolog.Logger
}

func New(colorService service.Color, log *zerolog.Logger) *UseCase {
	return &UseCase{
		colorService: colorService,
		log:          log,
	}
}
