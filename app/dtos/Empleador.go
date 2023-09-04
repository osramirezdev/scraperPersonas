package dtos

import "github.com/google/uuid"

type Empleador struct {
	ID          uuid.UUID `json:"id"`
	Patronal    string    `json:"patronal"`
	Empleador   string    `json:"empleador"`
	Estado      string    `json:"estado"`
	Aportes     int       `json:"aportes"`
	Vencimiento string    `json:"vencimiento"`
	Abonado     string    `json:"abonado"`
}
