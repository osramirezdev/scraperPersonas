package repository

import (
	"errors"
	"fmt"

	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/platform/database"
)

type ArchivoRepositoryI interface {
	Upsert(persona entities.Archivo) (entities.Archivo, error)
}

type archivoRepository struct {
	BaseRepo BaseRepositoryI[entities.Archivo]
}

var ArchivoRepository ArchivoRepositoryI

func init() {
	datab := &database.DBP
	ArchivoRepository = &archivoRepository{BaseRepo: &baseRepository[entities.Archivo]{Db: datab}}
}

// PersonRepo struct for queries from Persona entity.
func (query *archivoRepository) Upsert(archivo entities.Archivo) (entities.Archivo, error) {
	onconflict := "nombre"
	updates := []string{
		"full_direccion",
	}
	archivo, err := query.BaseRepo.CrearOnConflict(archivo, onconflict, updates)
	if err != nil {
		return entities.Archivo{}, errors.New("error creando archivo")
	}
	fmt.Println("archivo crearo", archivo)
	return archivo, nil
}
