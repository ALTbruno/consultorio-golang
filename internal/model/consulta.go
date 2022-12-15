package model

import "gorm.io/gorm"

type Consulta struct {
	gorm.Model
	IdDentista int    `json:"id_dentista"`
	IdPaciente int    `json:"id_pacient"`
	DataHora   string `json:"dataHora"`
	Descricao  string `json:"descricao"`
}

type Tabler interface {
	TableName() string
}
  
func (Consulta) TableName() string {
	return "consultas"
}
