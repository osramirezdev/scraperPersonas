package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/osramirezdev/scraperPersonas/app/dtos"
	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/app/services"
)

// Aqui abstraemos los metodos que queremos exponer
type PersonaControllerI interface {
	ObtenerDatosPersona(c *fiber.Ctx) error
}

// simulando constructor, y es solo privado, por que solo aqui podemos acceder a servicios
type personaController struct {
	personaService services.PersonaServiceI
}

// definimos como una variable global
var PersonaController PersonaControllerI

// y lo inicializamos
func init() {
	PersonaController = &personaController{
		personaService: services.PersonaService,
	}
}

// ObtenerDatos func realiza scraping de datos proveidos.
// @Description Se obtiene dato de persona por cedula, si no existe en base de datos se busca por scraping.
// @Summary busqueda datos de persona por cedula
// @Tags Personas
// @Accept json
// @Produce json
// @Param cedula path string true "Debe introducir un numero de cedula" minlength(6) maxlength(10)
// @Param tipo query string true "Debe introducir un tipo de busqueda, puede ser 'ruc', 'ips', 'funcionario'"
// @Success 200 {object} dtos.ResponseDto[int]
// @Failure 404 {object} dtos.ResponseDto404 "No se encontraron datos"
// @Failure 500 {object} dtos.ResponseDto500 "Internal server error"
// @Router /v1/persona/{cedula} [get]
func (p *personaController) ObtenerDatosPersona(c *fiber.Ctx) error {
	var respuesta dtos.ResponseDto[entities.Persona]
	cedula := c.Params("cedula")
	tipo := c.Params("tipo")
	persona, sErr := p.personaService.GetPersona(cedula, tipo)
	if sErr != nil {
		respuesta.Error = true
		respuesta.Data = nil
		respuesta.Msg = sErr.Error()
		return c.Status(fiber.StatusNotFound).JSON(respuesta)
	}
	respuesta.Count = 1
	respuesta.Data = []entities.Persona{persona}
	respuesta.Error = false
	respuesta.Msg = "Datos obtenidos"
	// Return status 200 OK.
	return c.JSON(respuesta)
}
