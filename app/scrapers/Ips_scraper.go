package scrapers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/osramirezdev/scraperPersonas/app/dtos"
)

type IpsScraperI interface {
	GetIpsFromScraping(ipsDto *dtos.IpsDto) ([]dtos.Asegurado, error)
}

type ipsScraper struct{}

var IpsScraper IpsScraperI = &ipsScraper{}

// obtenemos usuario.
func (i *ipsScraper) GetIpsFromScraping(ipsDto *dtos.IpsDto) ([]dtos.Asegurado, error) {
	datosSpliteados := ipsDto.Documentos

	var empleados []dtos.Asegurado
	for i, cedula := range datosSpliteados {
		fmt.Println(i)
		c := colly.NewCollector(
			colly.Debugger(&debug.LogDebugger{}),
		)

		var asegurados []dtos.Asegurado
		var empleadores []dtos.Empleador

		c.OnError(func(r *colly.Response, e error) {
			fmt.Println("Got this response:", r)
			fmt.Println("Got this error:", e)
		})

		// attach callbacks after login
		c.OnResponse(func(r *colly.Response) {
			fmt.Println("response received:", r.StatusCode)
			fmt.Println("Got response from:", r.Request.URL)
			fmt.Println("")
		})

		// capturamos la tabla que tiene un <input name="ekegir" />
		c.OnHTML("form > table:has(tbody):has(input[name='elegir'])", func(e *colly.HTMLElement) {

			item := dtos.Asegurado{
				Documento:            e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(2)"),
				Nombres:              e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(3)"),
				Apellidos:            e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(4)"),
				FechaNacimiento:      e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(5)"),
				Sexo:                 e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(6)"),
				Tipo:                 e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(7)"),
				BeneficiariosActivos: e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(8)"),
				Enrolado:             e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(9)"),
				Vencimiento:          e.ChildText("form > table:has(tbody):has(input[name='elegir']) tr td:nth-child(10)"),
			}
			asegurados = append(asegurados, item)

		})

		c.OnHTML("form > table:nth-of-type(3):has(td) tr[bgcolor='#e2e8f6']", func(e *colly.HTMLElement) {
			aporte, _ := strconv.Atoi(e.ChildText("td:nth-child(4)"))
			itemEmpleador := dtos.Empleador{
				Patronal:    e.ChildText("td:nth-child(1)"),
				Empleador:   e.ChildText("td:nth-child(2)"),
				Estado:      e.ChildText("td:nth-child(3)"),
				Aportes:     aporte,
				Vencimiento: e.ChildText("td:nth-child(5)"),
				Abonado:     e.ChildText("td:nth-child(6)"),
			}

			empleadores = append(empleadores, itemEmpleador)
			asegurados[0].Empleador = empleadores

		})

		c.Visit("https://servicios.ips.gov.py/consulta_asegurado/comprobacion_de_derecho_externo.php")

		err := c.Post("https://servicios.ips.gov.py/consulta_asegurado/comprobacion_de_derecho_externo.php", map[string]string{
			"nro_cic":   cedula,
			"recuperar": "Recuperar",
			"envio":     "ok",
		})
		if err != nil {
			return []dtos.Asegurado{}, errors.New(err.Error())
		}
		empleados = append(empleados, asegurados...)
		fmt.Println("cantidad empleados: ", len(empleados))
	}

	return empleados, nil
}
