package dto

type Dentista struct {
	Nome      string `json:"nome" gorm:"not null" validate:"required" example:"Fulano"`
	Sobrenome string `json:"sobrenome" gorm:"not null" validate:"required" example:"de Tal"`
	Matricula string `json:"matricula" gorm:"not null" validate:"required" example:"C123"`
}

type DentistaResponse struct {
	ID        uint   `json:"id" example:"3"`
	Nome      string `json:"nome" example:"Fulano"`
	Sobrenome string `json:"sobrenome" example:"de Tal"`
	Matricula string `json:"matricula" example:"C123"`
}
