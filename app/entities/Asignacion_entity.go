package entities

import (
	"gorm.io/gorm"
)

// User struct to describe User object.
type Asignacion struct {
	gorm.Model
	// funcionario
	Gasto               string `gorm:"type:varchar(150);" json:"gasto"`
	Estado              string `gorm:"type:varchar(150);" json:"estado"`
	Financiamiento      string `gorm:"type:varchar(150);" json:"financiamiento"`
	Linea               string `gorm:"type:varchar(150);" json:"linea"`
	Categoria           string `gorm:"type:varchar(150);" json:"categoria"`
	Cargo               string `gorm:"type:varchar(150);" json:"cargo"`
	Funcion             string `gorm:"type:varchar(150);" json:"funcion"`
	Presupuestado       int32  `gorm:"type:varchar(150);" json:"presupuestado"`
	Devengado           int32  `gorm:"type:varchar(150);" json:"devengado"`
	Movimiento          string `gorm:"type:varchar(150);" json:"movimiento"`
	Lugar               string `gorm:"type:varchar(150);" json:"lugar"`
	Actualizacion       string `gorm:"type:varchar(150);" json:"actualizacion"`
	Profesion           string `gorm:"type:varchar(150);" json:"profesion"`
	Correo              string `gorm:"type:varchar(150);" json:"correo"`
	Motivo              string `gorm:"type:varchar(150);" json:"motivo_movimiento"`
	FechaAdministrativo string `gorm:"type:varchar(150);" json:"fecha_administrativo"`
	Oficina             string `gorm:"type:varchar(150);" json:"oficina"`
	// docente
	Concepto            string `json:"concepto"`
	Institucion         string `json:"institucion"`
	AsignacionCategoria int32  `json:"asignacion_categoria"`
	Cantidad            int16  `json:"cantidad"`
	Asignacion          int32  `json:"asignacion"`
	Devuelto            string `json:"devuelto"`
	FuncionarioID       int
	Funcionario         Funcionario
}
