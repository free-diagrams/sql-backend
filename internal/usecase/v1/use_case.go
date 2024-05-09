package v1

import (
	"github.com/free-diagrams/sql-backend/internal/domain/infrastructure/repository"
	"github.com/free-diagrams/sql-backend/pkg/logger"
)

type UseCase struct {
	databaseRepository repository.DatabaseRepository
	userRepository     repository.UserRepository
	log                *logger.Logger
}

func New(databaseRepository repository.DatabaseRepository, userRepository repository.UserRepository,
	log *logger.Logger) *UseCase {
	return &UseCase{
		databaseRepository: databaseRepository,
		userRepository:     userRepository,
		log:                log,
	}
}
