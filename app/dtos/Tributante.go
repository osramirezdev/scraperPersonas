package dtos

import "time"

type Tributante struct {
	Id          uint       `db:"id" sql:"AUTO_INCREMENT" json:"id"`
	CreatedAt   string     `db:"created_at" json:"created_at"`
	UpdatedAt   string     `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
	Ruc         string     `db:"ruc" json:"ruc"`
	RazonSocial string     `db:"razon_social" json:"razon_social"`
	Estado      string     `db:"estado" json:"estado"`
	RucAnterior string     `db:"ruc_anterior" json:"ruc_anterior"`
}
