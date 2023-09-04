package dtos

import "github.com/google/uuid"

type Funcionario struct {
	ID               uuid.UUID    `json:"id"`
	Year             int16        `json:"anho"`
	Month            string       `json:"mes"`
	Nivel            string       `json:"nivel"`
	Entidad          string       `json:"entidad"`
	Oee              string       `json:"oee"`
	Documento        string       `json:"documento"`
	Nombres          string       `json:"nombres"`
	Apellidos        string       `json:"apellidos"`
	Presupuestado    int32        `json:"presupuestado"`
	Devengado        int32        `json:"devengado"`
	Sexo             string       `json:"sexo"`
	Estado           string       `json:"estado"`
	InitYear         int16        `json:"anho_inicio"`
	Discapacidad     string       `json:"discapacidad"`
	TipoDiscapacidad string       `json:"tipo_discapacidad"`
	Birthday         string       `json:"fecha_nacimiento"`
	Horario          string       `json:"horario"`
	Asignaciones     []Asignacion `json:"asignaciones"`
}
