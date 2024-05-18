package adapter

import (
	"github.com/free-diagrams/sql-backend/internal/infrastructure/repository/postgres"
	"github.com/free-diagrams/sql-backend/pkg/logger"
	"github.com/free-diagrams/sql-backend/pkg/storage/storagesqlx"
)

type Adapter struct {
	db  storagesqlx.Queryer
	log *logger.Logger
}

func New(db *postgres.Postgres, log *logger.Logger) *Adapter {
	return &Adapter{
		db:  db,
		log: log,
	}
}
