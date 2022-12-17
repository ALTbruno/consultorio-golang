package model

import "gorm.io/gorm"

type Paciente struct {
	gorm.Model
	Nome           string   `json:"nome" gorm:"not null" validate:"required"`
	Sobrenome      string   `json:"sobrenome" gorm:"not null" validate:"required"`
	RG             string   `json:"rg" gorm:"not null" validate:"required"`
	DataDeCadastro string   `json:"dataDeCadastro"`
	Consulta       Consulta `json:"-"`
}

func (Paciente) TableName() string {
	return "pacientes"
}
