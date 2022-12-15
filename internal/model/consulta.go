package model

type Consulta struct {
	IdDentista	  int    `json:"id_dentista"`
	IdPaciente  int    `json:"id_pacient"`
	DataHora  string `json:"dataHora"`
	Descricao string `json:"descricao"`
}
