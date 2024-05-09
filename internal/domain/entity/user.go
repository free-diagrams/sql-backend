package entity

import "github.com/google/uuid"

type User struct {
	ID        *string
	Username  string
	LastName  string
	FirstName string
}

func (e *User) GenerateID() {
	generatedID := uuid.New().String()
	e.ID = &generatedID
}
