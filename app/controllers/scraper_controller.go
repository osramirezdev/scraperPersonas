package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/osramirezdev/scraperPersonas/app/dtos"
	scrapers "github.com/osramirezdev/scraperPersonas/app/scrapers"
	"github.com/osramirezdev/scraperPersonas/pkg/utils"
)

// ScrapRuc func realiza scraping de datos proveidos.
// @Description Permite consultar si ruc existen por cedula o nombre.
// @Summary scraping de documentos o nombres
// @Tags Scraper
// @Accept json
// @Produce json
// @Param ruc body dtos.RucDto true "Nombres es campo opcional"
// @Success 200 {object} dtos.ResponseDto[int]
// @Failure 404 {object} dtos.ResponseDto404 "No se encontro ruc"
// @Failure 500 {object} dtos.ResponseDto500 "Internal server error"
// @Router /v1/scraper/obtener/ruc/ [post]
func ScrapRuc(c *fiber.Ctx) error {
	var respuesta dtos.ResponseDto[dtos.Tributante]
	// Create new ruc struct
	ruc := dtos.RucDto{}
	fmt.Println(ruc)
	// Create a new validator for a body dto.
	validate := utils.NewValidator()

	// Check, if received JSON data is valid.
	if errV := c.BodyParser(&ruc); errV != nil {
		fmt.Println("error parse", errV)
		respuesta.Error = true
		respuesta.Data = nil
		respuesta.Msg = errV.Error()
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(respuesta)
	}

	// Validate dto fields.
	if errD := validate.Struct(ruc); errD != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(errD),
		})
	}

	scrap, sErr := scrapers.RucScraper.GetRucFromScraping(&ruc)
	if sErr != nil {
		respuesta.Error = true
		respuesta.Data = nil
		respuesta.Msg = sErr.Error()
		return c.Status(fiber.StatusNotFound).JSON(respuesta)
	}
	respuesta.Count = len(scrap)
	respuesta.Data = scrap
	respuesta.Error = false
	respuesta.Msg = "Datos obtenidos"
	// Return status 200 OK.
	return c.JSON(respuesta)
}

// ScrapIps func realiza scraping de cedulas proveidas para obtener datos de asegurados.
// @Description Hacer scraping de ips.
// @Summary Permite consultar datos de asegiurados por cedula
// @Tags Scraper
// @Accept json
// @Produce json
// @Param ips body dtos.IpsDto true "Dto para consultar datos"
// @Success 200 {object} dtos.ResponseDto[dtos.Asegurado]
// @Failure 404 {object} dtos.ResponseDto404 "No se encontro asegurado"
// @Failure 500 {object} dtos.ResponseDto500
// @Router /v1/scraper/obtener/ips/ [post]
func ScrapIps(c *fiber.Ctx) error {
	var respuesta dtos.ResponseDto[dtos.Asegurado]

	ips := &dtos.IpsDto{}

	// Create a new validator for a body dto.
	validate := utils.NewValidator()

	// Check, if received JSON data is valid.
	if errV := c.BodyParser(&ips); errV != nil {
		fmt.Println("error parse", errV)
		respuesta.Error = true
		respuesta.Data = nil
		respuesta.Msg = errV.Error()
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(respuesta)
	}

	// Validate dto fields.
	if errD := validate.Struct(ips); errD != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(errD),
		})
	}

	scrap, sErr := scrapers.IpsScraper.GetIpsFromScraping(ips)
	if sErr != nil {
		respuesta.Error = true
		respuesta.Data = nil
		respuesta.Msg = sErr.Error()
		return c.Status(fiber.StatusNotFound).JSON(respuesta)
	}
	respuesta.Count = len(scrap)
	respuesta.Data = scrap
	respuesta.Error = false
	respuesta.Msg = "Datos obtenidos"
	// Return status 200 OK.
	return c.JSON(respuesta)
}

// ScrapFuncionarios realiza scraping de cedulas proveidas para obtener datos de funcionarios.
// @Description Validad si una cedula pertenece a un funcionario publico.
// @Summary Permite consultar si datos de funcionarios existen por cedula, mes y anho
// @Tags Scraper
// @Accept json
// @Produce json
// @Param funcionario body dtos.FuncionarioDto true "Dto para consultar datos"
// @Success 200 {object} dtos.ResponseDto[dtos.Funcionario]
// @Failure 404 {object} dtos.ResponseDto404 "No se encontro asegurado"
// @Failure 500 {object} dtos.ResponseDto500
// @Router /v1/scraper/obtener/funcionarios/ [post]
func ScrapFuncionarios(c *fiber.Ctx) error {
	var respuesta dtos.ResponseDto[dtos.Funcionario]

	data := &dtos.FuncionarioDto{}

	// Create a new validator for a body dto.
	validate := utils.NewValidator()

	// Check, if received JSON data is valid.
	if errV := c.BodyParser(&data); errV != nil {
		fmt.Println("error parse", errV)
		respuesta.Error = true
		respuesta.Data = nil
		respuesta.Msg = errV.Error()
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(respuesta)
	}

	// Validate dto fields.
	if errD := validate.Struct(data); errD != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(errD),
		})
	}

	scrap, sErr := scrapers.FuncionarioScraper.GetFuncionarioFromScraping(data)
	if sErr != nil {
		respuesta.Error = true
		respuesta.Data = nil
		respuesta.Msg = sErr.Error()
		return c.Status(fiber.StatusNotFound).JSON(respuesta)
	}
	respuesta.Count = len(scrap)
	respuesta.Data = scrap
	respuesta.Error = false
	respuesta.Msg = "Datos obtenidos"
	// Return status 200 OK.
	return c.JSON(respuesta)
}

// // ScrapUploadCSV Espera obtener un archivo csv con parametros dados para realizar busqueda de datos de los mismos.
// // @Description Alzar un archivo csv para que el scraper realice busquedas.
// // @Summary Permite realizar busquedas de datos a partir de un archivo proveido
// // @Tags Scraper
// // @Accept json
// // @Produce json
// // @Param archivo formData file true "Alzar archivo csv"
// // @Success 200 {object} dtos.ResponseDto[string]
// // @Failure 404 {object} dtos.ResponseDto404 "No se encontro asegurado"
// // @Failure 500 {object} dtos.ResponseDto500
// // @Router /v1/scraper/upload/ [post]
// func ScrapUploadCSV(c *fiber.Ctx) error {
// 	var respuesta dtos.ResponseDto[string]
// 	fileForm, err := c.FormFile("archivo")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		respuesta.Error = true
// 		respuesta.Data = nil
// 		respuesta.Msg = err.Error()
// 		return c.Status(fiber.StatusInternalServerError).JSON(respuesta)
// 	}
// 	file, errFile := fileForm.Open()
// 	if errFile != nil {
// 		fmt.Println("error abrir archivo", errFile)
// 		respuesta.Error = true
// 		respuesta.Data = nil
// 		respuesta.Msg = errFile.Error()
// 		return c.Status(fiber.StatusInternalServerError).JSON(respuesta)
// 	}
// 	columnas := []string{
// 		"cedula", "nombre", "apellido",
// 	}
// 	validateCSV := utils.NewCSVValidator(columnas)
// 	errorLeyendoArchiv := validateCSV.Validate(file)
// 	if errorLeyendoArchiv != nil {
// 		fmt.Println("error file1", errorLeyendoArchiv)
// 		respuesta.Error = true
// 		respuesta.Data = nil
// 		respuesta.Msg = errorLeyendoArchiv.Error()
// 		return c.Status(fiber.StatusInternalServerError).JSON(respuesta)
// 	}
// 	dbFile, dbErr := services.FileService.SaveToDB(fileForm, fileForm.Filename)
// 	if dbErr != nil {
// 		fmt.Println("error file2", dbErr)
// 		respuesta.Error = true
// 		respuesta.Data = nil
// 		respuesta.Msg = dbErr.Error()
// 		return c.Status(fiber.StatusInternalServerError).JSON(respuesta)
// 	}
// 	saved := dbFile.Nombre
// 	respuesta.Error = false
// 	respuesta.Data = []string{saved}
// 	respuesta.Msg = "Archivo recibido correctamente"
// 	// Return status 200 OK.
// 	return c.JSON(respuesta)
// }
