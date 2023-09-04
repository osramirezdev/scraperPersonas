package scrapers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/stealth"
	"github.com/google/uuid"
	"github.com/osramirezdev/scraperPersonas/app/dtos"
)

type FuncionarioScraperI interface {
	GetFuncionarioFromScraping(docenteDto *dtos.FuncionarioDto) ([]dtos.Funcionario, error)
}

type funcionarioScraper struct{}

var FuncionarioScraper FuncionarioScraperI = &funcionarioScraper{}

func (f *funcionarioScraper) GetFuncionarioFromScraping(docenteDto *dtos.FuncionarioDto) ([]dtos.Funcionario, error) {
	var mensajeImpreso bool
	query := docenteDto.Search

	browser := rod.New().MustConnect()

	defer func() {
		// Cierra el navegador al finalizar
		err := browser.Close()
		if err != nil {
			log.Fatal("err ", err)
		}
	}()

	var funcs []dtos.Funcionario
	page := stealth.MustPage(browser)

	page = page.MustNavigate("https://datos.sfp.gov.py/data/funcionarios").MustWaitStable()

	for _, q := range query {

		// Le ordenamos al navegador que espere a que la tabla este vidsible
		page.MustElement("table#DataTables_Table_0:has(tbody):has(tr)").MustWaitLoad().MustWaitVisible()
		valueInput := page.MustElement("input#anho_filter").MustText()
		valueMes := page.MustElement(`select[id="mes_filter"]`).MustText()
		d := fmt.Sprintf(`() => document.querySelector('input[id="documento_filter"]').value = '%s'`, q.Documento)

		documento := page.MustEval(d).Str()
		page.MustElement(`input[id="documento_filter"]`).MustType(input.Enter)

		// buscar todos los meses
		fmt.Println("buscar todos los meses...")
		page.MustEval(`() => document.querySelector("select#mes_filter").value = ""`)
		page.MustEval(`() => document.querySelector("select#mes_filter").focus()`)
		page.MustElement("select#mes_filter").MustWaitVisible().MustWaitLoad().MustType(input.ArrowDown)
		page.MustElement("select#mes_filter").MustWaitVisible().MustWaitLoad().MustType(input.ArrowDown)
		page.MustElement("select#mes_filter").MustWaitVisible().MustWaitLoad().MustType(input.ArrowUp)
		page.MustElement("select#mes_filter").MustWaitVisible().MustWaitLoad().MustType(input.ArrowUp)

		fmt.Println("anho: ", valueInput, "mes: ", valueMes, documento)
		page.MustElement(`div#DataTables_Table_0_processing[style="display: block;"]`).MustWaitLoad().MustWaitVisible()
		ajaxLoader := page.MustElements(`div#DataTables_Table_0_processing[style="display: block;"]`)
		fmt.Println("ajaxLoader: ", ajaxLoader)
		// // while este el elemento loading
		for len(ajaxLoader) > 0 {
			if !mensajeImpreso {
				fmt.Println("tabla esta siendo consultada...")
				mensajeImpreso = true
			}
			ajaxLoader = page.MustElements(`div#DataTables_Table_0_processing[style="display: block;"]`)
		}
		fmt.Println("ya se poblo la tabla, continuamos...")

		fmt.Println("ya se poblo la tabla, continuamos...")
		notFound := page.MustElements(`tbody:has([class="dataTables_empty"])`)
		if len(notFound) > 0 {
			fmt.Println("No se han encontrado resultados.")
			return funcs, nil
		}
		// actualizamos contenido pagina
		page.MustElement("table#DataTables_Table_0:has(tbody):has(tr)").MustWaitLoad().MustWaitVisible()
		// page.MustElement("ul.pagination li a:not(#DataTables_Table_0_first):not(#DataTables_Table_0_previous):not(#DataTables_Table_0_last):not(#DataTables_Table_0_next)").MustWaitVisible()
		page.MustElement(`#DataTables_Table_0_last`).MustWaitStable().MustWaitVisible()
		paginas := page.MustElements("ul.pagination li:nth-last-child(3)")
		// paginas := page.MustElements("ul.pagination li")
		fmt.Println("paginas: ", len(paginas), paginas[0].MustElement("a").MustText())
		numeroPaginas, _ := strconv.Atoi(paginas[0].MustElement("a").MustText())
		if len(paginas) > 0 {
			for i := 0; i < numeroPaginas; i++ {
				NavigatePage(page, &funcs)
				siguiente := page.MustElements("li:not(.disabled) > #DataTables_Table_0_next")
				fmt.Println("siguiente: ", siguiente, i)
				if len(siguiente) > 0 {
					fmt.Println("primero debe dar click")
					page.MustElement("li:not(.disabled) > #DataTables_Table_0_next").MustWaitVisible().MustClick()
					page.MustElement(`div#DataTables_Table_0_processing[style="display: block;"]`).MustWaitLoad().MustWaitVisible()
					ajaxLoader := page.MustElements(`div#DataTables_Table_0_processing[style="display: block;"]`)
					fmt.Println("ajaxLoader: ", ajaxLoader)
					// // while este el elemento loading
					for len(ajaxLoader) > 0 {
						if !mensajeImpreso {
							fmt.Println("tabla esta siendo consultada...")
							mensajeImpreso = true
						}
						ajaxLoader = page.MustElements(`div#DataTables_Table_0_processing[style="display: block;"]`)
					}
				}
			}
			fmt.Println("fin iter cantidad funcionarios: ", len(funcs))
		}
		// anhos := []string{"2023", "2022", "2021", "2020", "2019", "2018", "2017", "2016", "2015"}
		// for i := 0; i < len(anhos); i++ {

		// }
		// for _, anho := range anhos {
		// 	fmt.Println("Current anho: ", anho)

		// 	fmt.Println("clickea antes de terminar")
		// 	page.MustElement("input#anho_filter").MustType(input.ArrowDown)
		// }

		// esperamos que la pagina se cargue completamente
		page.MustWaitLoad()
	}

	return funcs, nil
}

func NavigatePage(page *rod.Page, funcionarios *[]dtos.Funcionario) {
	// Utiliza selectores CSS para obtener los datos de la tabla
	// obtenemos el select con su atributo name, y seleccionamos en este caso el maximo numero
	rows := page.MustElements("tbody tr")
	// for len(rows) <= 50 {
	// 	fmt.Println("Sigue siendo menor a 1")
	// 	rows = page.MustElements("tbody tr")
	// }
	i := 1
	currenIndex := 0
	for _, row := range rows {
		year, _ := strconv.Atoi(row.MustElement("td:nth-child(2)").MustText())

		// The dot is a metacharactere. It matches single character. To use it as a real dot, you have to escape it with \
		monto := regexp.MustCompile(`\.`).ReplaceAllString(row.MustElement("td:nth-child(10)").MustText(), "")
		deve := regexp.MustCompile(`\.`).ReplaceAllString(row.MustElement("td:nth-child(11)").MustText(), "")
		presupuestado, _ := strconv.Atoi(monto)
		devengado, _ := strconv.Atoi(deve)
		ingreso, _ := strconv.Atoi(row.MustElement("td:nth-child(14)").MustText())

		itemFuncionario := dtos.Funcionario{
			ID:               uuid.New(),
			Year:             int16(year),
			Month:            row.MustElement("td:nth-child(3)").MustText(),
			Nivel:            row.MustElement("td:nth-child(4)").MustText(),
			Entidad:          row.MustElement("td:nth-child(5)").MustText(),
			Oee:              row.MustElement("td:nth-child(6)").MustText(),
			Documento:        row.MustElement("td:nth-child(7)").MustText(),
			Nombres:          row.MustElement("td:nth-child(8)").MustText(),
			Apellidos:        row.MustElement("td:nth-child(9)").MustText(),
			Presupuestado:    int32(presupuestado),
			Devengado:        int32(devengado),
			Sexo:             row.MustElement("td:nth-child(12)").MustText(),
			Estado:           row.MustElement("td:nth-child(13)").MustText(),
			InitYear:         int16(ingreso),
			Discapacidad:     row.MustElement("td:nth-child(15)").MustText(),
			TipoDiscapacidad: row.MustElement("td:nth-child(16)").MustText(),
			Birthday:         row.MustElement("td:nth-child(17)").MustText(),
			Horario:          row.MustElement("td:nth-child(18)").MustText(),
		}

		// capturamos la primera columna de la fila que estamos
		index := fmt.Sprintf(`tr:nth-child(%v) td:nth-child(1)`, i)
		el := page.MustElement(index)
		el.MustClick()
		var mensajeImpreso bool
		page.MustElement(`[id$="_processing"]:not([style="display: none;"])`).MustWaitLoad().MustWaitVisible()
		ajaxLoaderSub := page.MustElements(`[id$="_processing"]:not([style="display: none;"])`)
		// // while este el elemento loading
		for len(ajaxLoaderSub) > 0 {
			if !mensajeImpreso {
				fmt.Println("subtabla esta siendo consultada...")
				mensajeImpreso = true
			}
			ajaxLoaderSub = page.MustElements(`[id$="_processing"]:not([style="display: none;"])`)
		}
		fmt.Println("despues: ")
		findIndex := fmt.Sprintf(`table#DataTables_Table_0_%v:has(tbody)`, currenIndex)
		fmt.Println("stuck: ", index, findIndex)
		page.MustElement(findIndex).MustWaitLoad().MustWaitVisible()
		findIndex = findIndex + " tbody tr"
		fmt.Println("buscando tr", findIndex)
		tr := page.MustElement(findIndex).MustWaitVisible()
		fmt.Println("encontrado tr", tr)
		rowsTwo := page.MustElements(findIndex)
		for _, rowTwo := range rowsTwo {
			// The dot is a metacharactere. It matches single character. To use it as a real dot, you have to escape it with \
			monto := regexp.MustCompile(`\.`).ReplaceAllString(rowTwo.MustElement("td:nth-child(8)").MustText(), "")
			deve := regexp.MustCompile(`\.`).ReplaceAllString(rowTwo.MustElement("td:nth-child(9)").MustText(), "")
			presupuestado, _ := strconv.Atoi(monto)
			devengado, _ := strconv.Atoi(deve)
			itemAsignacion := dtos.Asignacion{
				ID:                  uuid.New(),
				Gasto:               rowTwo.MustElement("td:nth-child(1)").MustText(),
				Estado:              rowTwo.MustElement("td:nth-child(2)").MustText(),
				Financiamiento:      rowTwo.MustElement("td:nth-child(3)").MustText(),
				Linea:               rowTwo.MustElement("td:nth-child(4)").MustText(),
				Categoria:           rowTwo.MustElement("td:nth-child(5)").MustText(),
				Cargo:               rowTwo.MustElement("td:nth-child(6)").MustText(),
				Funcion:             rowTwo.MustElement("td:nth-child(7)").MustText(),
				Presupuestado:       int32(presupuestado),
				Devengado:           int32(devengado),
				Movimiento:          rowTwo.MustElement("td:nth-child(10)").MustText(),
				Lugar:               rowTwo.MustElement("td:nth-child(11)").MustText(),
				Actualizacion:       rowTwo.MustElement("td:nth-child(12)").MustText(),
				Profesion:           rowTwo.MustElement("td:nth-child(13)").MustText(),
				Correo:              rowTwo.MustElement("td:nth-child(14)").MustText(),
				Motivo:              rowTwo.MustElement("td:nth-child(15)").MustText(),
				FechaAdministrativo: rowTwo.MustElement("td:nth-child(16)").MustText(),
				Oficina:             rowTwo.MustElement("td:nth-child(17)").MustText(),
			}
			itemFuncionario.Asignaciones = append(itemFuncionario.Asignaciones, itemAsignacion)
		}
		*funcionarios = append(*funcionarios, itemFuncionario)
		i = i + 1
		currenIndex = currenIndex + 1
		el.MustClick()
	}
}
