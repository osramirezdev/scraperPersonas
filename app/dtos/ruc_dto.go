package dtos

// Ruc dto: para solicitar consulta al scraper.
// se envian arreglos de cedulas o nombres
type RucDto struct {
	Documentos []string `json:"documentos" validate:"required,lte=500,dive"`
	// optional fields
	Nombres []string `json:"nombres" validate:"lte=500,dive"`
}
