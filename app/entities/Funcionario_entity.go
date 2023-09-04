package entities

import (
	"gorm.io/gorm"
)

// User struct to describe User object.
type Funcionario struct {
	gorm.Model
	Anho                string `gorm:"type:varchar(6);" json:"anho" validate:"lte=6"`
	Mes                 string `gorm:"type:varchar(15);" json:"mes" validate:"lte=15"`
	Nivel               string `gorm:"type:varchar(150);" json:"nivel" validate:"lte=150"`
	Entidad             string `gorm:"type:varchar(150);" json:"entidad" validate:"lte=150"`
	Oee                 string `gorm:"type:varchar(150);" json:"oee" validate:"lte=150"`
	Presupuestado       string `gorm:"type:varchar(10);" json:"presupuestado" validate:"lte=10"`
	Devengado           string `gorm:"type:varchar(10);" json:"devengado" validate:"lte=10"`
	Estado              string `gorm:"type:varchar(50);" json:"estado" validate:"lte=50"`
	AnhoIngreso         string `gorm:"type:varchar(6);" json:"anho_ingreso" validate:"lte=6"`
	Discapacidad        string `gorm:"type:varchar(6);" json:"discapacidad"`
	TipoDiscapacidad    string `gorm:"type:varchar(6);" json:"tipo_discapacidad"`
	Horario             string `gorm:"type:varchar(150);" json:"horario"`
	ObjetoGasto         string `gorm:"type:varchar(150);" json:"objeto_gasto"`
	Antiguedad          string `gorm:"type:varchar(150);" json:"antiguedad"`
	NumeroMatriculacion string `gorm:"type:varchar(150);" json:"numero_matriculacion"`
	Asignacion          string `gorm:"type:varchar(10);" json:"asignacion"`
	PersonaID           int
	Persona             Persona
}
