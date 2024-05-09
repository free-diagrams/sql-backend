package userusecase

import (
	"github.com/free-diagrams/sql-backend/internal/domain/infrastructure"
	"github.com/free-diagrams/sql-backend/pkg/logger"
)

type UseCase struct {
	userRepository infrastructure.UserRepository
	log            *logger.Logger
}

func New(userRepository infrastructure.UserRepository, log *logger.Logger) *UseCase {
	return &UseCase{
		userRepository: userRepository,
		log:            log,
	}
}
