package entity

import domainentity "github.com/free-diagrams/sql-backend/internal/domain/entity"

type Database struct {
	ID   *string `db:"id"`
	Name string  `db:"name"`
}

func (e *Database) ToDomain() *domainentity.Database {
	return &domainentity.Database{
		ID:   e.ID,
		Name: e.Name,
	}
}
