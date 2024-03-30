package usecase

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/usecase"
	"github.com/pkg/errors"
)

func (u *UseCase) GetColors(ctx context.Context) ([]usecase.GetColorsResponseItem, error) {
	colors, err := u.colorService.GetColors(ctx)
	if err != nil {
		return make([]usecase.GetColorsResponseItem, 0), errors.Wrap(err, "failed to get colors")
	}

	response := make([]usecase.GetColorsResponseItem, len(colors))
	for i, color := range colors {
		response[i] = usecase.GetColorsResponseItem{
			ID:      color.ID,
			Name:    color.Name,
			HexCode: color.HexCode,
		}
	}

	return response, nil
}
