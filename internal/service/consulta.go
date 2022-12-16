package service

import (
	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/repository"
)

func CadastrarConsulta(consulta model.Consulta) model.Consulta {
	c := repository.CadastrarConsulta(consulta)
	return c
}

func BuscarConsultaPorID(id int) model.Consulta {
	consulta := repository.BuscarConsultaPorID(id)
	return consulta
}

func DeletarConsultaPorID(id int) {
	repository.DeletarConsultaPorID(id)
}
