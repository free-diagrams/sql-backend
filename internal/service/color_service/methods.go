package color_service

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/pkg/errors"
)

func (s *ColorService) GetColors(ctx context.Context) ([]entity.Color, error) {
	colors, err := s.repo.GetColors(ctx)
	if err != nil {
		return make([]entity.Color, 0), errors.Wrap(err, "failed to get colors")
	}

	lang := ctx.Value("lang").(string)
	for i := range colors {
		newColorName, err := s.localizer.TryLocalize(lang, colors[i].Name)
		if err != nil {
			s.log.Debug().Ctx(ctx).Str("colorName", colors[i].Name).Msg("Failed to localize color")
			newColorName = s.localizer.English(colors[i].Name)
		}
		colors[i].Name = newColorName
	}

	return colors, nil
}
