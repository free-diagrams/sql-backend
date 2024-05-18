package adapter

import (
	"context"
	domainentity "github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/free-diagrams/sql-backend/internal/domain/errors"
	"github.com/free-diagrams/sql-backend/internal/infrastructure/repository/postgres/entity"
	"github.com/free-diagrams/sql-backend/internal/util"
	"github.com/free-diagrams/sql-backend/pkg/errs"
)

func (a *Adapter) GetAllDatabases(ctx context.Context) ([]domainentity.Database, error) {
	queryer := util.ExtractTx(ctx)
	if queryer == nil {
		queryer = a.db
	}

	query := `
		SELECT id, name
		FROM databases
	`

	databases := make([]entity.Database, 0)
	err := queryer.SelectContext(ctx, &databases, query)
	if err != nil {
		return nil, errs.WrapErrorError(errors.InternalDatabase, err)
	}
	if len(databases) == 0 {
		return nil, errors.DatabaseNotFound
	}

	domainDatabases := make([]domainentity.Database, len(databases))
	for i, database := range databases {
		domainDatabases[i] = *database.ToDomain()
	}

	return domainDatabases, nil
}
