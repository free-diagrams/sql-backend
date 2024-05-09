package v1

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
)

// TODO: Нормально реализовать этот метод
func (u *UseCase) GetUserByToken(context context.Context, accessToken string) (*entity.User, error) {
	id := "kolesnikovda1"
	return &entity.User{
		ID:        &id,
		Username:  "kolesnikovda1",
		LastName:  "Колесников",
		FirstName: "Дмитрий",
	}, nil
}
