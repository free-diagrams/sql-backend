package database_repository

import (
	"context"
	domainentity "github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/free-diagrams/sql-backend/internal/infrastructure/repository/entity"
	"github.com/free-diagrams/sql-backend/pkg/errs"
	"github.com/free-diagrams/sql-backend/pkg/storage/storagesql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (r *Repository) GetAllDatabases(ctx context.Context) ([]domainentity.Database, error) {
	return r.getAllDatabases(r.db, ctx)
}

func (r *Repository) GetAllDatabasesTx(tx *sqlx.Tx, ctx context.Context) ([]domainentity.Database, error) {
	return r.getAllDatabases(tx, ctx)
}

func (r *Repository) getAllDatabases(queryer storagesql.Queryer, ctx context.Context) ([]domainentity.Database, error) {
	query := `
		SELECT id, name
		FROM databases
	`

	databases := make([]entity.Database, 0)
	err := queryer.SelectContext(ctx, &databases, query)
	if err != nil {
		r.log.Debug().Err(err).Msg("failed to execute query")
		return nil, errors.Wrap(err, errs.InternalDatabase.Error())
	}
	if len(databases) == 0 {
		return nil, errs.DatabaseNotFound
	}

	domainDatabases := make([]domainentity.Database, len(databases))
	for i, database := range databases {
		domainDatabases[i] = *database.ToDomain()
	}

	return domainDatabases, nil
}
