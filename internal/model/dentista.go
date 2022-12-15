package model

import "gorm.io/gorm"

type Dentista struct {
	gorm.Model
	Nome      string   `json:"nome"`
	Sobrenome string   `json:"sobrenome"`
	Matricula string   `json:"matricula"`
	Consulta  Consulta `json:"-"`
}

func (Dentista) TableName() string {
	return "dentistas"
}
