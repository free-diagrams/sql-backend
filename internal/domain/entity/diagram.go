package entity

import (
	"github.com/google/uuid"
)

type Diagram struct {
	ID        *string
	Tables    []Table
	Relations []Relation
	Name      string
	IsPublic  bool
	Comment   *string
}

func (e *Diagram) GenerateID() {
	generatedID := uuid.New().String()
	e.ID = &generatedID
}
