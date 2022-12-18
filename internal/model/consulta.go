package model

import "gorm.io/gorm"

type Consulta struct {
	gorm.Model
	DentistaID uint   `json:"id_dentista" gorm:"not null" validate:"required"`
	PacienteID uint   `json:"id_paciente" gorm:"not null" validate:"required"`
	DataHora   string `json:"dataHora" gorm:"not null" validate:"required"`
	Descricao  string `json:"descricao" gorm:"not null" validate:"required"`
}

type Tabler interface {
	TableName() string
}

func (Consulta) TableName() string {
	return "consultas"
}
