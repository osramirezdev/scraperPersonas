package services

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/osramirezdev/scraperPersonas/app/dtos"
	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/app/repository"
	scrapers "github.com/osramirezdev/scraperPersonas/app/scrapers"
)

type PersonaServiceI interface {
	GetPersona(cedula string, tipo string) (entities.Persona, error)
}

type personaService struct {
	PersonaRepo        repository.PersonaRepositoryI
	FuncionarioRepo    repository.FuncionarioRepositoryI
	AsignacionRepo     repository.AsignacionesRepositoryI
	IpsRepo            repository.IpsRepositoryI
	RucScraper         scrapers.RucScraperI
	IpsScraper         scrapers.IpsScraperI
	FuncionarioScraper scrapers.FuncionarioScraperI
}

var PersonaService PersonaServiceI

func init() {
	PersonaService = &personaService{
		PersonaRepo:        repository.PersonaRepository,
		FuncionarioRepo:    repository.FuncionarioRepository,
		AsignacionRepo:     repository.AsignacionRepository,
		IpsRepo:            repository.IpsRepository,
		RucScraper:         scrapers.RucScraper,
		IpsScraper:         scrapers.IpsScraper,
		FuncionarioScraper: scrapers.FuncionarioScraper,
	}
}

// obtenemos usuario.
func (repo *personaService) GetPersona(cedula string, tipo string) (entities.Persona, error) {
	personaCreated := entities.Persona{
		Documento: cedula,
	}
	persona, err := repo.PersonaRepo.Upsert(personaCreated)
	if err != nil {
		fmt.Println("Error encontrando persona: ", err)
		return entities.Persona{}, errors.New(err.Error())
	}

	fmt.Println("persona encontrada jaja", persona)
	repo.personaScraper(&persona, cedula, tipo)
	fmt.Println("no deberia llegar aqui hasta tener persona, ruc al menos", persona)
	return persona, nil
}

func (repo *personaService) personaScraper(p *entities.Persona, cedula string, tipo string) {
	fmt.Println("id [persona]: ", p.ID, p.Model.ID)
	rucDto := &dtos.RucDto{
		Documentos: []string{cedula},
	}
	ipsDto := &dtos.IpsDto{
		Documentos: []string{cedula},
	}
	funcionarioDto := &dtos.FuncionarioDto{
		Search: []dtos.Search{
			{
				Documento: cedula,
				Tipo:      tipo,
			},
		},
	}
	var ipsEntity []entities.Ips
	var funcionarioEntity []entities.Funcionario
	var asignacionesEntity []entities.Asignacion

	// busca datos ruc
	datosRuc, _ := repo.RucScraper.GetRucFromScraping(rucDto)
	if reflect.ValueOf(datosRuc).IsNil() {
		fmt.Println("Ruc Vacio, se continua: ")
	} else {
		p.Documento = cedula
		p.Ruc = strings.ReplaceAll(datosRuc[0].Ruc, " ", "")
		p.RazonSocial = datosRuc[0].RazonSocial
		p.Nombre = strings.Split(datosRuc[0].RazonSocial, ",")[1]
		p.Apellido = strings.Split(datosRuc[0].RazonSocial, ",")[0]
	}

	// busca datos ips
	datosIps, _ := repo.IpsScraper.GetIpsFromScraping(ipsDto)
	if reflect.ValueOf(datosIps).IsNil() {
		fmt.Println("Ips Vacio, se continua: ")
	} else {
		p.Documento = cedula
		p.Sexo = datosIps[0].Sexo
		p.FechaNacimiento = datosIps[0].FechaNacimiento
		for _, i := range datosIps {
			if len(i.Empleador) > 0 {
				for _, z := range i.Empleador {
					var ipsResponse entities.Ips
					ipsResponse.Persona.ID = p.Model.ID
					ipsResponse.Abonado = z.Abonado
					ipsResponse.Aportes = z.Aportes
					ipsResponse.BeneficiariosActivos = i.BeneficiariosActivos
					ipsResponse.Enrolado = i.Enrolado
					ipsResponse.Vencimiento = z.Vencimiento
					ipsResponse.Estado = z.Estado
					ipsResponse.Empleador = z.Empleador
					ipsResponse.Patronal = z.Patronal
					ipsEntity = append(ipsEntity, ipsResponse)
				}
			}
		}
	}
	// busca datos ruc

	fmt.Println(funcionarioDto)
	datosFuncionario, _ := repo.FuncionarioScraper.GetFuncionarioFromScraping(funcionarioDto)
	// se busca funcionario
	if reflect.ValueOf(datosFuncionario).IsNil() {
		fmt.Println("Funcionario Vacio, se continua: ")
	} else {
		p.Documento = cedula
		p.Nombre = datosFuncionario[0].Nombres
		p.Apellido = datosFuncionario[0].Apellidos
		p.Sexo = datosFuncionario[0].Sexo
		p.FechaNacimiento = datosFuncionario[0].Birthday
		for _, f := range datosFuncionario {
			var fResponse entities.Funcionario
			fResponse.Persona.ID = p.Model.ID
			fResponse.Anho = strconv.Itoa(int(f.Year))
			fResponse.Mes = f.Month
			fResponse.Nivel = f.Nivel
			fResponse.Entidad = f.Entidad
			fResponse.Oee = f.Oee
			fResponse.Presupuestado = strconv.Itoa(int(f.Presupuestado))
			fResponse.Devengado = strconv.Itoa(int(f.Devengado))
			fResponse.Estado = f.Estado
			fResponse.AnhoIngreso = strconv.Itoa(int(f.InitYear))
			fResponse.Discapacidad = f.Discapacidad
			fResponse.TipoDiscapacidad = f.TipoDiscapacidad
			fResponse.Horario = f.Horario
			if len(f.Asignaciones) > 0 {
				for _, z := range f.Asignaciones {
					var aResponse entities.Asignacion
					aResponse.Gasto = z.Gasto
					aResponse.Estado = z.Estado
					aResponse.Financiamiento = z.Financiamiento
					aResponse.Linea = z.Linea
					aResponse.Categoria = z.Categoria
					aResponse.Cargo = z.Cargo
					aResponse.Funcion = z.Funcion
					aResponse.Presupuestado = z.Presupuestado
					aResponse.Devengado = z.Devengado
					aResponse.Movimiento = z.Movimiento
					aResponse.Lugar = z.Lugar
					aResponse.Actualizacion = z.Actualizacion
					aResponse.Profesion = z.Profesion
					aResponse.Correo = z.Correo
					aResponse.Motivo = z.Motivo
					aResponse.FechaAdministrativo = z.FechaAdministrativo
					aResponse.Oficina = z.Oficina
					asignacionesEntity = append(asignacionesEntity, aResponse)
				}
			}
			funcionarioEntity = append(funcionarioEntity, fResponse)
		}
	}
	if len(ipsEntity) > 0 {
		ipsCreated, _ := repo.IpsRepo.CreateMany(ipsEntity)
		fmt.Println("ips: ", ipsCreated)
	}
	if len(funcionarioEntity) > 0 {
		funcionarioCreated, _ := repo.FuncionarioRepo.CreateMany(funcionarioEntity)
		fmt.Println("funcionarios: ", funcionarioCreated)
		fmt.Println("funcionario id: ", funcionarioCreated[0].ID)
		funcionarioEntity = funcionarioCreated
	}
	if len(asignacionesEntity) > 0 {
		for _, a := range asignacionesEntity {
			a.Funcionario.ID = funcionarioEntity[0].Model.ID
		}
		asignacionCreated, _ := repo.AsignacionRepo.CreateMany(asignacionesEntity)
		fmt.Println("asignaciones: ", asignacionCreated)
	}
	person, _ := repo.PersonaRepo.Upsert(*p)
	*p = person
	fmt.Println("id persona: ", p, len(ipsEntity), len(funcionarioEntity), funcionarioEntity)
}
