package dto

type Consulta struct {
	DentistaID uint   `json:"id_dentista" gorm:"not null" validate:"required"`
	PacienteID uint   `json:"id_paciente" gorm:"not null" validate:"required"`
	DataHora   string `json:"dataHora" gorm:"not null" validate:"required" example:"23-02-2021 14:30"`
	Descricao  string `json:"descricao" gorm:"not null" validate:"required" example:"Limpeza e extração"`
}
