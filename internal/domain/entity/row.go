package entity

import "github.com/google/uuid"

type Row struct {
	ID           *string
	DataType     DataType
	ColorRGB     string
	LogicalName  string
	PhysicalName string
	IsPrimaryKey bool
	IsNullable   bool
	IsUnique     bool
	DefaultValue *string
	Comment      *string
}

func (e *Row) GenerateID() {
	generatedID := uuid.New().String()
	e.ID = &generatedID
}
