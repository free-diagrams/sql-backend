package usecase

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
)

type DatabaseUseCase interface {
	GetAllDatabases(ctx context.Context) ([]entity.Database, error)
}
