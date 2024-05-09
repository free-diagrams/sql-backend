package database_repository

import (
	"github.com/free-diagrams/sql-backend/internal/infrastructure/repository/postgres"
	"github.com/free-diagrams/sql-backend/pkg/logger"
)

type Repository struct {
	db  *postgres.Postgres
	log *logger.Logger
}

func New(db *postgres.Postgres, log *logger.Logger) *Repository {
	return &Repository{
		db:  db,
		log: log,
	}
}
