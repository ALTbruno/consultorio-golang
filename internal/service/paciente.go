package service

import (
	"time"

	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/repository"
)

func CadastrarPaciente(paciente model.Paciente) model.Paciente {
	hoje := time.Now().Format("02-01-2006")
	paciente.DataDeCadastro = hoje
	p := repository.CadastrarPaciente(paciente)
	return p
}

func BuscarPacientePorID(id int) model.Paciente {
	paciente := repository.BuscarPacientePorID(id)
	return paciente
}

func DeletarPacientePorID(id int) {
	repository.DeletarPacientePorID(id)
}
