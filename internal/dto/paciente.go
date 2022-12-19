package dto

type Paciente struct {
	Nome      string `json:"nome" gorm:"not null" validate:"required" example:"Ciclano"`
	Sobrenome string `json:"sobrenome" gorm:"not null" validate:"required" example:"Silva"`
	RG        string `json:"rg" gorm:"not null" validate:"required" example:"789456"`
}

type PacienteResponse struct {
	ID             uint   `json:"id" example:"19"`
	Nome           string `json:"nome" example:"Ciclano"`
	Sobrenome      string `json:"sobrenome" example:"Silva"`
	RG             string `json:"rg" example:"789456"`
	DataDeCadastro string `json:"dataDeCadastro" example:"23-03-2021"`
}
