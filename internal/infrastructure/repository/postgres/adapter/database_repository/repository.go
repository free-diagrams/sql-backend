package database_repository

import (
	"github.com/free-diagrams/sql-backend/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db  *sqlx.DB
	log *logger.Logger
}

func New(db *sqlx.DB, log *logger.Logger) *Repository {
	return &Repository{
		db:  db,
		log: log,
	}
}
