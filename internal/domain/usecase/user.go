package usecase

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
)

type UserUseCase interface {
	GetUserByToken(context context.Context, accessToken string) (*entity.User, error)
}
