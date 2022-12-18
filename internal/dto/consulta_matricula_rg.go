package dto

type ConsultaMatriculaRG struct {
	MatriculaDentista string `json:"matricula_dentista" validate:"required"`
	RGPaciente        string `json:"rg_paciente" validate:"required"`
	DataHora          string `json:"dataHora" validate:"required"`
	Descricao         string `json:"descricao" validate:"required"`
}
