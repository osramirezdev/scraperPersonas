package repository

import (
	"errors"
	"fmt"

	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/platform/database"
	"gorm.io/gorm/clause"
)

type IpsRepositoryI interface {
	Upsert(persona entities.Ips) (entities.Ips, error)
	CreateMany(ips []entities.Ips) ([]entities.Ips, error)
}

type ipsRepository struct {
	BaseRepo BaseRepositoryI[entities.Ips]
}

var IpsRepository IpsRepositoryI

func init() {
	datab := &database.DBP
	IpsRepository = &ipsRepository{BaseRepo: &baseRepository[entities.Ips]{Db: datab}}
}

// PersonRepo struct for queries from Persona entity.
func (query *ipsRepository) Upsert(ips entities.Ips) (entities.Ips, error) {
	onconflict := "patronal, "
	updates := []string{
		"empleador",
		"vencimiento",
		"estado",
		"aportes",
	}
	person, err := query.BaseRepo.CrearOnConflict(ips, onconflict, updates)
	if err != nil {
		return entities.Ips{}, errors.New("error creando ips")
	}
	fmt.Println("ips creado", person)
	return ips, nil
}

func (query *ipsRepository) CreateMany(ips []entities.Ips) ([]entities.Ips, error) {
	cols := []clause.Column{
		{
			Name: "patronal",
		},
		{
			Name: "persona_id",
		},
	}
	updates := []string{
		"empleador",
		"vencimiento",
		"estado",
		"aportes",
	}
	person, err := query.BaseRepo.CrearManyConflict(ips, cols, updates)
	if err != nil {
		return []entities.Ips{}, errors.New("error creando ips")
	}
	fmt.Println("ips creado", person)
	return ips, nil
}
