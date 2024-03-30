package color_repository

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/free-diagrams/sql-backend/pkg/errs"
	"github.com/pkg/errors"
)

func (r *ColorRepository) GetColors(ctx context.Context) ([]entity.Color, error) {
	query := `
		SELECT *
		FROM colors
	`

	colors := make([]entity.Color, 0)
	err := r.queryer.SelectContext(ctx, &colors, query)
	if err != nil {
		r.log.Debug().Err(err).Msg("Failed to get colors from database")
		return make([]entity.Color, 0), errors.Wrap(errs.InternalDatabase, "failed to get colors from database")
	}

	return colors, nil
}
