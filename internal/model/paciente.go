package model

type Paciente struct {
	Nome           string `json:"nome"`
	Sobrenome      string `json:"sobrenome"`
	RG             string `json:"matricula"`
	DataDeCadastro string `json:"dataDeCadastro"`
}
