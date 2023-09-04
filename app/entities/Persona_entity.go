package entities

import (
	"gorm.io/gorm"
)

// User struct to describe User object.
type Persona struct {
	gorm.Model
	Documento       string `gorm:"type:varchar(25); index:idx_documento,unique" json:"documento" validate:"required,lte=50"`
	Ruc             string `gorm:"type:varchar(25); index:idx_ruc" json:"ruc" validate:"required,lte=25"`
	RazonSocial     string `gorm:"type:varchar(150);" json:"razon_social"`
	Estado          string `gorm:"type:varchar(150);" json:"estado"`
	RucAnterior     string `gorm:"type:varchar(25);" json:"ruc_anterior"`
	FechaNacimiento string `gorm:"type:varchar(25)" validate:"required,lte=25" json:"fecha_nacimiento"`
	Nombre          string `gorm:"type:varchar(70)" json:"nombre" validate:"lte=70"`
	Apellido        string `gorm:"type:varchar(70)" json:"apellido" validate:"lte=70"`
	Sexo            string `gorm:"type:varchar(20)" validate:"lte=20" json:"sexo"`
}
