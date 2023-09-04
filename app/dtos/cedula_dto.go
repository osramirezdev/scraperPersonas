package dtos

// Ruc dto: para solicitar consulta al scraper.
// se envian arreglos de cedulas o nombres
type Cedula struct {
	Cedula string `json:"cedula" validate:"required,lte=12"`
}
