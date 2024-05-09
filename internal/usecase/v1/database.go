package v1

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/pkg/errors"
)

func (u *UseCase) GetAllDatabases(ctx context.Context) ([]entity.Database, error) {
	databases, err := u.databaseRepository.GetAllDatabases(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all databases")
	}

	return databases, nil
}
