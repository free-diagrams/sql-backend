package postgres

import (
	"fmt"
	"github.com/free-diagrams/sql-backend/internal/config"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"strings"
	"time"
)

type Postgres struct {
	*sqlx.DB
}

func New(cfg config.DBConfig) (*Postgres, error) {
	connectionStringBuilder := strings.Builder{}
	connectionStringBuilder.WriteString(fmt.Sprintf("postgresql://%s", cfg.Username))
	connectionStringBuilder.WriteString(fmt.Sprintf(":%s", cfg.Password))
	connectionStringBuilder.WriteString(fmt.Sprintf("@%s", cfg.Host))
	connectionStringBuilder.WriteString(fmt.Sprintf(":%d", cfg.Port))
	connectionStringBuilder.WriteString(fmt.Sprintf("/%s", cfg.Name))
	connectionStringBuilder.WriteString("?sslmode=disable")
	connectionString := connectionStringBuilder.String()

	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	pgdb := Postgres{
		db,
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime * time.Minute)

	goose.SetTableName("migrations")

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err := goose.Up(db.DB, cfg.MigrationPath); err != nil {
		return nil, err
	}

	return &pgdb, nil
}
