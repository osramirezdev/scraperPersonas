package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/osramirezdev/scraperPersonas/app/dtos"
)

type TributanteRepository struct {
	db *sqlx.DB
}

// instanciamos tributante repository.
func NewTributanteRepository(db *sqlx.DB) *TributanteRepository {
	return &TributanteRepository{db: db}
}

// obtenemos usuario.
func (cr *TributanteRepository) Get(id int64) *dtos.Tributante {
	var tributante dtos.Tributante
	return &tributante
}
