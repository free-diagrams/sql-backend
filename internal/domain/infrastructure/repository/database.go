package repository

import (
	"context"
	domainentity "github.com/free-diagrams/sql-backend/internal/domain/entity"
)

type DatabaseRepository interface {
	GetAllDatabases(ctx context.Context) ([]domainentity.Database, error)
}
