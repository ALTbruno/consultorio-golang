package dto

type Dentista struct {
	Nome      string `json:"nome" gorm:"not null" validate:"required" example:"Fulano"`
	Sobrenome string `json:"sobrenome" gorm:"not null" validate:"required" example:"de Tal"`
	Matricula string `json:"matricula" gorm:"not null" validate:"required" example:"C123"`
}
