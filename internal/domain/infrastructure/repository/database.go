package repository

import (
	"context"
	domainentity "github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type DatabaseRepository interface {
	GetAllDatabases(ctx context.Context) ([]domainentity.Database, error)
	GetAllDatabasesTx(tx *sqlx.Tx, ctx context.Context) ([]domainentity.Database, error)
}
