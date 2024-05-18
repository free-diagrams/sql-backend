package util

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
	"github.com/free-diagrams/sql-backend/pkg/storage/storagesqlx"
	"github.com/jmoiron/sqlx"
)

func ExtractTx(ctx context.Context) storagesqlx.Queryer {
	tx, ok := ctx.Value(entity.ContextTransaction).(*sqlx.Tx)
	if !ok {
		return nil
	}

	return tx
}
