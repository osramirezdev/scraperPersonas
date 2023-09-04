package repository

import (
	"errors"
	"fmt"

	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/platform/database"
	"gorm.io/gorm/clause"
)

type AsignacionesRepositoryI interface {
	Upsert(funcionario entities.Asignacion) (entities.Asignacion, error)
	CreateMany(funcionario []entities.Asignacion) ([]entities.Asignacion, error)
}

type asignacionesRepository struct {
	BaseRepo BaseRepositoryI[entities.Asignacion]
}

var AsignacionRepository AsignacionesRepositoryI

func init() {
	datab := &database.DBP
	AsignacionRepository = &asignacionesRepository{BaseRepo: &baseRepository[entities.Asignacion]{Db: datab}}
}

// PersonRepo struct for queries from Persona entity.
func (query *asignacionesRepository) Upsert(funcionario entities.Asignacion) (entities.Asignacion, error) {
	onconflict := "patronal, "
	updates := []string{
		"empleador",
		"vencimiento",
		"estado",
		"aportes",
	}
	person, err := query.BaseRepo.CrearOnConflict(funcionario, onconflict, updates)
	if err != nil {
		return entities.Asignacion{}, errors.New("error creando funcionario")
	}
	fmt.Println("funcionario creado", person)
	return funcionario, nil
}

func (query *asignacionesRepository) CreateMany(funcionario []entities.Asignacion) ([]entities.Asignacion, error) {
	cols := []clause.Column{
		{
			Name: "funciona",
		},
		{
			Name: "anho",
		},
		{
			Name: "funcionario_id",
		},
	}
	updates := []string{
		"gasto",
		"estado",
		"financiamiento",
		"linea",
		"categoria",
		"cargo",
		"funcion",
		"presupuestado",
		"devengado",
		"movimiento",
		"lugar",
		"actualizacion",
		"profesion",
		"correo",
		"motivo",
		"fecha_administrativo",
		"oficina",
		// docente
		"concepto",
		"institucion",
		"asignacion_categoria",
		"cantidad",
		"devuelto",
	}
	person, err := query.BaseRepo.CrearManyConflict(funcionario, cols, updates)
	if err != nil {
		return []entities.Asignacion{}, errors.New("error creando funcionario")
	}
	fmt.Println("funcionario creado", person)
	return funcionario, nil
}
