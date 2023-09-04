package services

import (
	"errors"
	"fmt"

	"github.com/osramirezdev/scraperPersonas/app/dtos"
	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/app/repository"
	scrapers "github.com/osramirezdev/scraperPersonas/app/scrapers"
)

type FuncionarioServiceI interface {
	GetFuncionario(cedula *string, mes string, nmes string, anho string) ([]entities.Funcionario, error)
}

type funcionarioService struct {
	FuncionarioRepo    repository.FuncionarioRepositoryI
	FuncionarioScraper scrapers.FuncionarioScraperI
}

var FuncionarioService FuncionarioServiceI

func init() {
	FuncionarioService = &funcionarioService{
		FuncionarioRepo: repository.FuncionarioRepository,
	}
}

// obtenemos usuario.
func (repo *funcionarioService) GetFuncionario(cedula *string, mes string, nmes string, anho string) ([]entities.Funcionario, error) {
	// personaCreated := entities.Funcionario{}
	persona, err := repo.FuncionarioScraper.GetFuncionarioFromScraping(&dtos.FuncionarioDto{})
	fmt.Println("persona: ", persona, " error: ", err)
	if err != nil {
		if err.Error() != "record not found" {
			fmt.Println("Error encontrando persona: ", err)
			return []entities.Funcionario{}, errors.New(err.Error())
		}
	}

	fmt.Println("no deberia llegar aqui hasta tener persona, ruc al menos", persona)
	return []entities.Funcionario{}, nil
}
