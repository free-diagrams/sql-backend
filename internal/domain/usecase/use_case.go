package usecase

import (
	"context"
	"github.com/free-diagrams/sql-backend/internal/domain/entity"
)

type UseCase interface {
	DatabaseUseCase
	UserUseCase
}

type DatabaseUseCase interface {
	GetAllDatabases(ctx context.Context) ([]entity.Database, error)
}

type UserUseCase interface {
	GetUserByToken(context context.Context, accessToken string) (*entity.User, error)
}
