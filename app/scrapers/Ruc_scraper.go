package scrapers

import (
	"fmt"
	"log"

	"github.com/go-rod/rod"
	"github.com/osramirezdev/scraperPersonas/app/dtos"
)

type RucScraperI interface {
	GetRucFromScraping(rucDto *dtos.RucDto) ([]dtos.Tributante, error)
}

type rucScraper struct{}

var RucScraper RucScraperI

func init() {
	RucScraper = &rucScraper{}
}

// obtenemos usuario.
func (r *rucScraper) GetRucFromScraping(rucDto *dtos.RucDto) ([]dtos.Tributante, error) {
	datosSpliteados := append(rucDto.Documentos, rucDto.Nombres...)

	browser := rod.New().MustConnect()

	defer func() {
		// Cierra el navegador al finalizar
		err := browser.Close()
		if err != nil {
			log.Fatal("err ", err)
		}
	}()

	page := browser.MustPage("https://www.ruc.com.py/").MustWaitLoad()
	var tributantes []dtos.Tributante
	for i, cadena := range datosSpliteados {

		inputci := page.MustElement(`[id="txt_buscar"]`)

		searchBtn := page.MustElement(`[id="btn_buscar"]`).MustWaitEnabled().MustWaitVisible()

		inputci.MustSelectAllText().MustInput("")

		inputci.MustInput(cadena)

		searchBtn.MustClick()

		getRuc(page, &tributantes)

		fmt.Println("termino vuelta: ", i)

	}

	return tributantes, nil
}

func getRuc(page *rod.Page, tributantes *[]dtos.Tributante) {
	var mensajeImpreso bool
	buscando := true
	page.MustElement(`[id="base_ruc"]`).MustWaitLoad().MustWaitVisible()

	// page.MustElement(`table[id="base_ruc"] tbody tr:has(td[colspan="2"])`).MustWaitLoad().MustWaitVisible()
	// while este el elemento loading
	for buscando {
		if !mensajeImpreso {
			fmt.Println("tabla esta siendo consultada...")
			mensajeImpreso = true
		}
		ajaxLoader := page.MustElements(`table[id="base_ruc"] tbody tr:has(td[colspan="2"]) td[width="100%"]`)
		if len(ajaxLoader) == 0 {
			buscando = false
		}
	}

	notFound := page.MustElements(`table[id="base_ruc"] tbody tr:has(td[colspan="2"])`)
	success := page.MustElements(`table[id="base_ruc"] tbody tr`)
	if len(notFound) > 0 {
		return
	} else if len(success) > 0 {
		success = page.MustElements(`table[id="base_ruc"] tbody tr`)
		fmt.Println("cantidad filas: ", len(success))
		for i, row := range success {
			fmt.Println(i)
			itemSet := dtos.Tributante{
				Ruc:         row.MustElement(`td:nth-child(1)`).MustText(),
				RazonSocial: row.MustElement(`td:nth-child(2)`).MustText(),
				Estado:      "desconocido",
			}
			*tributantes = append(*tributantes, itemSet)
		}
	} else {
		return
	}
}
