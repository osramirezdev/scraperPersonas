package dtos

import "github.com/google/uuid"

type Docente struct {
	ID                  uuid.UUID `json:"id"`
	Anho                int16     `json:"anho"`
	Mes                 string    `json:"mes"`
	Documento           string    `json:"documento"`
	NombreCompleto      string    `json:"nombre_completo"`
	Estado              string    `json:"estado"`
	ObjetoGasto         string    `json:"objeto_gasto"`
	Antiguedad          string    `json:"antiguedad"`
	NumeroMatriculacion int32     `json:"numero_matriculacion"`
	Asignacion          int32     `json:"asignacion"`
	Sexo                string    `json:"sexo"`
	Detalles            []Detalle `json:"detalles"`
}
