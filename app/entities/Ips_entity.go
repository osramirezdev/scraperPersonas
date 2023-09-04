package entities

import "gorm.io/gorm"

type Ips struct {
	gorm.Model
	Tipo                 string `gorm:"type:varchar(25)" json:"tipo"`
	BeneficiariosActivos string `gorm:"type:varchar(25)" json:"beneficiarios_activos"`
	Enrolado             string `gorm:"type:varchar(10)" json:"enrolado"`
	Vencimiento          string `gorm:"type:varchar(25)" json:"vencimiento"`
	Patronal             string `gorm:"type:varchar(70)" json:"patronal"`
	Empleador            string `gorm:"type:varchar(130)" json:"empleador"`
	Estado               string `gorm:"type:varchar(25)" json:"estado"`
	Aportes              int    `gorm:"type:varchar(25)" json:"aportes"`
	Abonado              string `gorm:"type:varchar(25)" json:"abonado"`
	PersonaID            int
	Persona              Persona
}
