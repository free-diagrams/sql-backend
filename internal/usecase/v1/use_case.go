package v1

import (
	"github.com/free-diagrams/sql-backend/internal/domain/infrastructure/repository"
	"github.com/free-diagrams/sql-backend/pkg/logger"
	"github.com/free-diagrams/sql-backend/pkg/storage/storagesqlx/txmanager"
)

type UseCase struct {
	txManager          txmanager.TxManager
	databaseRepository repository.DatabaseRepository
	userRepository     repository.UserRepository
	log                *logger.Logger
}

func New(txManager txmanager.TxManager, databaseRepository repository.DatabaseRepository, userRepository repository.UserRepository,
	log *logger.Logger) *UseCase {
	return &UseCase{
		txManager:          txManager,
		databaseRepository: databaseRepository,
		userRepository:     userRepository,
		log:                log,
	}
}
