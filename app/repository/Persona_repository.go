package repository

import (
	"errors"
	"fmt"

	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/platform/database"
)

type PersonaRepositoryI interface {
	Upsert(persona entities.Persona) (entities.Persona, error)
	Get(cedula string) (entities.Persona, error)
	FindOne(cedula string) (entities.Persona, error)
}

type personaRepository struct {
	BaseRepo BaseRepositoryI[entities.Persona]
}

var PersonaRepository PersonaRepositoryI

func init() {
	datab := &database.DBP
	PersonaRepository = &personaRepository{BaseRepo: &baseRepository[entities.Persona]{Db: datab}}
}

// PersonRepo struct for queries from Persona entity.
func (query *personaRepository) Upsert(persona entities.Persona) (entities.Persona, error) {
	onconflict := "documento"
	updates := []string{
		"ruc",
		"nombre",
		"apellido",
	}
	person, err := query.BaseRepo.CrearOnConflict(persona, onconflict, updates)
	if err != nil {
		return entities.Persona{}, errors.New(err.Error())
	}
	fmt.Println("persona creada", person)
	return person, nil
}

func (query *personaRepository) Get(cedula string) (entities.Persona, error) {
	person, err := query.BaseRepo.GetByAttribute(entities.Persona{}, "documento", cedula)
	if err != nil {
		fmt.Println("error buscando persona: ", err)
		return entities.Persona{}, errors.New(err.Error())
	}
	fmt.Println("persona encontrada", person.ID, person.Model.ID, person)
	return person, nil

}

func (query *personaRepository) FindOne(cedula string) (entities.Persona, error) {
	person, err := query.BaseRepo.FindOne(entities.Persona{}, "documento", cedula)
	if err != nil {
		fmt.Println("error buscando persona: ", err)
		return entities.Persona{}, errors.New(err.Error())
	}
	fmt.Println("persona encontrada", person.ID, person.Model.ID, person)
	return person, nil
}
