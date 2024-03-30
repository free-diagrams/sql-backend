package color_service

import (
	"github.com/free-diagrams/sql-backend/internal/domain/repository"
	"github.com/free-diagrams/sql-backend/pkg/loclzr"
	"github.com/rs/zerolog"
)

type ColorService struct {
	repo      repository.Color
	localizer *loclzr.Localizer
	log       *zerolog.Logger
}

func New(repo repository.Color, localizer *loclzr.Localizer, log *zerolog.Logger) *ColorService {
	return &ColorService{
		repo:      repo,
		localizer: localizer,
		log:       log,
	}
}
