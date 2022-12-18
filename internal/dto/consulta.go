package dto

type ConsultaBody struct {
	DentistaID uint   `json:"id_dentista" gorm:"not null" validate:"required" example:"3"`
	PacienteID uint   `json:"id_paciente" gorm:"not null" validate:"required" example:"19"`
	DataHora   string `json:"dataHora" gorm:"not null" validate:"required" example:"23-02-2021 14:30"`
	Descricao  string `json:"descricao" gorm:"not null" validate:"required" example:"Limpeza e extração"`
}

type ConsultaBodyMatriculaRG struct {
	MatriculaDentista string `json:"matricula_dentista" gorm:"not null" validate:"required" example:"C123"`
	RGPaciente        string `json:"rg_paciente" gorm:"not null" validate:"required" example:"789456"`
	DataHora          string `json:"dataHora" gorm:"not null" validate:"required" example:"23-02-2021 14:30"`
	Descricao         string `json:"descricao" gorm:"not null" validate:"required" example:"Limpeza e extração"`
}

type ConsultaResponse struct {
	ID         uint   `json:"id" gorm:"not null" validate:"required" example:"50"`
	DentistaID uint   `json:"id_dentista" gorm:"not null" validate:"required" example:"3"`
	PacienteID uint   `json:"id_paciente" gorm:"not null" validate:"required"  example:"19"`
	DataHora   string `json:"dataHora" gorm:"not null" validate:"required" example:"23-02-2021 14:30"`
	Descricao  string `json:"descricao" gorm:"not null" validate:"required"  example:"Limpeza e extração"`
}
