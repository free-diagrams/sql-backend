package repository

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
)

type Color interface {
	GetColors(ctx context.Context) ([]entity.Color, error)
}
