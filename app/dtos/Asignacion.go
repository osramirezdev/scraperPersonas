package dtos

import "github.com/google/uuid"

type Asignacion struct {
	ID                  uuid.UUID `json:"id"`
	Gasto               string    `json:"gasto"`
	Estado              string    `json:"estado"`
	Financiamiento      string    `json:"financiamiento"`
	Linea               string    `json:"linea"`
	Categoria           string    `json:"categoria"`
	Cargo               string    `json:"cargo"`
	Funcion             string    `json:"funcion"`
	Presupuestado       int32     `json:"presupuestado"`
	Devengado           int32     `json:"devengado"`
	Movimiento          string    `json:"movimiento"`
	Lugar               string    `json:"lugar"`
	Actualizacion       string    `json:"actualizacion"`
	Profesion           string    `json:"profesion"`
	Correo              string    `json:"correo"`
	Motivo              string    `json:"motivo_movimiento"`
	FechaAdministrativo string    `json:"fecha_administrativo"`
	Oficina             string    `json:"oficina"`
}
