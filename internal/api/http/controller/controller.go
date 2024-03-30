package controller

import (
	"github.com/free-diagrams/sql-backend/internal/domain/usecase"
	"github.com/free-diagrams/sql-backend/pkg/loclzr"
	"github.com/rs/zerolog"
)

type Controller struct {
	useCase   usecase.UseCase
	localizer *loclzr.Localizer
	log       *zerolog.Logger
}

func New(useCase usecase.UseCase, localizer *loclzr.Localizer, log *zerolog.Logger) *Controller {
	return &Controller{
		useCase:   useCase,
		localizer: localizer,
		log:       log,
	}
}
