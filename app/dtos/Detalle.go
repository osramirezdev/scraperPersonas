package dtos

import "github.com/google/uuid"

type Detalle struct {
	ID                  uuid.UUID `db:"id" sql:"AUTO_INCREMENT" json:"id"`
	Concepto            string    `json:"concepto"`
	Institucion         string    `json:"institucion"`
	Cargo               string    `json:"cargo"`
	Documento           string    `json:"documento"`
	Categoria           string    `json:"categoria"`
	AsignacionCategoria int32     `json:"asignacion_categoria"`
	Cantidad            int16     `json:"cantidad"`
	Asigancion          int32     `json:"asignacion"`
	Devuelto            string    `json:"devuelto"`
}
