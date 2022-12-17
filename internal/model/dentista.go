package model

import "gorm.io/gorm"

type Dentista struct {
	gorm.Model
	Nome      string   `json:"nome" gorm:"not null" validate:"required"`
	Sobrenome string   `json:"sobrenome" gorm:"not null" validate:"required"`
	Matricula string   `json:"matricula" gorm:"not null" validate:"required"`
	Consulta  Consulta `json:"-"`
}

func (Dentista) TableName() string {
	return "dentistas"
}
