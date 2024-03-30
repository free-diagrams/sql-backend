package color_repository

import (
	"github.com/free-diagrams/sql-backend/pkg/storage/storagesql"
	"github.com/rs/zerolog"
)

type ColorRepository struct {
	queryer storagesql.Queryer
	log     *zerolog.Logger
}

func New(queryer storagesql.Queryer, log *zerolog.Logger) *ColorRepository {
	return &ColorRepository{
		queryer: queryer,
		log:     log,
	}
}
