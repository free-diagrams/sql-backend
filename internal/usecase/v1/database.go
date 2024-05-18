package v1

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	domainentity "github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/pkg/errors"
)

func (u *UseCase) GetAllDatabases(ctx context.Context) ([]entity.Database, error) {
	var databases []domainentity.Database

	err := u.txManager.Transactional(ctx, func(ctx context.Context) error {
		var err error

		databases, err = u.databaseRepository.GetAllDatabases(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to get all databases")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return databases, nil
}
