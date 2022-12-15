package model

import "gorm.io/gorm"

type Paciente struct {
	gorm.Model
	Nome           string   `json:"nome"`
	Sobrenome      string   `json:"sobrenome"`
	RG             string   `json:"rg"`
	DataDeCadastro string   `json:"dataDeCadastro"`
	Consulta       Consulta `json:"-"`
}

func (Paciente) TableName() string {
	return "pacientes"
}
