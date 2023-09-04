package dtos

// Ruc dto: para solicitar consulta al scraper.
// se envian arreglos de cedulas o nombres
type IpsDto struct {
	Documentos []string `json:"documentos" validate:"required,lte=500,dive"`
}
