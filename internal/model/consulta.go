package model

import "gorm.io/gorm"

type Consulta struct {
	gorm.Model
	DentistaID uint   `json:"id_dentista"`
	PacienteID uint   `json:"id_paciente"`
	DataHora   string `json:"dataHora"`
	Descricao  string `json:"descricao"`
}

type Tabler interface {
	TableName() string
}

func (Consulta) TableName() string {
	return "consultas"
}
