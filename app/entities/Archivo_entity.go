package entities

import (
	"gorm.io/gorm"
)

// User struct to describe User object.
type Archivo struct {
	gorm.Model
	Nombre        string `gorm:"type:varchar(50); unique" json:"nombre" validate:"required,lte=50" `
	FullDireccion string `gorm:"type:varchar(70); unique" json:"full_direccion" validate:"required,lte=70"`
	Leido         bool   `gorm:"type:boolean;default:false" json:"ruc" validate:"required,lte=70"`
}
