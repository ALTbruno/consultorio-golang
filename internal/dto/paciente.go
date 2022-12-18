package dto

type Paciente struct {
	Nome      string `json:"nome" gorm:"not null" validate:"required" example:"Ciclano"`
	Sobrenome string `json:"sobrenome" gorm:"not null" validate:"required" example:"Silva"`
	RG        string `json:"rg" gorm:"not null" validate:"required" example:"789456"`
}
