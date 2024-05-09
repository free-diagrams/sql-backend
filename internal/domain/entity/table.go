package entity

import "github.com/google/uuid"

type Table struct {
	ID           *string
	Rows         []Row
	ColorRGB     string
	LogicalName  string
	PhysicalName string
	Comment      *string
}

func (e *Table) GenerateID() {
	generatedID := uuid.New().String()
	e.ID = &generatedID
}
