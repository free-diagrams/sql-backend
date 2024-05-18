package txmanager

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	internalerr "github.com/free-diagrams/sql-backend/internal/domain/errors"
	"github.com/free-diagrams/sql-backend/internal/infrastructure/repository/postgres"
	"github.com/free-diagrams/sql-backend/pkg/errs"
	"github.com/free-diagrams/sql-backend/pkg/logger"
)

type TxManager struct {
	db  *postgres.Postgres
	log *logger.Logger
}

func New(db *postgres.Postgres, log *logger.Logger) *TxManager {
	return &TxManager{
		db:  db,
		log: log,
	}
}

func (m *TxManager) Transactional(ctx context.Context, fn func(context.Context) error) error {
	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return errs.WrapErrorError(internalerr.InternalDatabase, err)
	}

	txContext := context.WithValue(ctx, entity.ContextTransaction, tx)

	err = fn(txContext)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			m.log.Error().Err(errRollback).Msg("failed to rollback transaction")
		}
		return errs.WrapErrorError(internalerr.InternalDatabase, err)
	}

	err = tx.Commit()
	if err != nil {
		return errs.WrapErrorError(internalerr.InternalDatabase, err)
	}
	return nil
}
