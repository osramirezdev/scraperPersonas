package dtos

import "github.com/google/uuid"

type Asegurado struct {
	ID                   uuid.UUID   `json:"id"`
	Documento            string      `json:"documento"`
	Nombres              string      `json:"nombres"`
	Apellidos            string      `json:"apellidos"`
	FechaNacimiento      string      `json:"fecha_nacimiento"`
	Sexo                 string      `json:"sexo"`
	Tipo                 string      `json:"tipo"`
	BeneficiariosActivos string      `json:"beneficiarios_activos"`
	Enrolado             string      `json:"enrolado"`
	Vencimiento          string      `json:"vencimiento"`
	Empleador            []Empleador `json:"empleadores"`
}
