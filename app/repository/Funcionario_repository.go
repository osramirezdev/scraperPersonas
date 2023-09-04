package repository

import (
	"errors"
	"fmt"

	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/platform/database"
	"gorm.io/gorm/clause"
)

type FuncionarioRepositoryI interface {
	Upsert(funcionario entities.Funcionario) (entities.Funcionario, error)
	CreateMany(funcionario []entities.Funcionario) ([]entities.Funcionario, error)
}

type funcionarioRepository struct {
	BaseRepo BaseRepositoryI[entities.Funcionario]
}

var FuncionarioRepository FuncionarioRepositoryI

func init() {
	datab := &database.DBP
	FuncionarioRepository = &funcionarioRepository{BaseRepo: &baseRepository[entities.Funcionario]{Db: datab}}
}

// PersonRepo struct for queries from Persona entity.
func (query *funcionarioRepository) Upsert(funcionario entities.Funcionario) (entities.Funcionario, error) {
	onconflict := "patronal"
	updates := []string{
		"empleador",
		"vencimiento",
		"estado",
		"aportes",
	}
	person, err := query.BaseRepo.CrearOnConflict(funcionario, onconflict, updates)
	if err != nil {
		return entities.Funcionario{}, errors.New("error creando funcionario")
	}
	fmt.Println("funcionario creado", person)
	return funcionario, nil
}

func (query *funcionarioRepository) CreateMany(funcionario []entities.Funcionario) ([]entities.Funcionario, error) {
	cols := []clause.Column{
		{
			Name: "mes",
		},
		{
			Name: "anho",
		},
		{
			Name: "persona_id",
		},
	}
	updates := []string{
		"nivel",
		"entidad",
		"oee",
		"presupuestado",
		"devengado",
		"estado",
		"anho_ingreso",
		"discapacidad",
		"tipo_discapacidad",
		"horario",
		"objeto_gasto",
		"antiguedad",
		"numero_matriculacion",
	}
	person, err := query.BaseRepo.CrearManyConflict(funcionario, cols, updates)
	if err != nil {
		return []entities.Funcionario{}, errors.New("error creando funcionario")
	}
	fmt.Println("funcionario creado", person)
	return funcionario, nil
}
