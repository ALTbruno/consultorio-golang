package repository

import (
	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/ALTbruno/consultorio-golang/internal/model"
)

func CadastrarPaciente(paciente model.Paciente) model.Paciente {
	db.DB.Create(&paciente)
	return paciente
}

func BuscarPacientePorID(id int) model.Paciente {
	var paciente model.Paciente
	db.DB.First(&paciente, id)
	return paciente
}

func AtualizarPacientePorID(id int) model.Paciente {
	paciente := BuscarPacientePorID(id)
	db.DB.First(&paciente, id)
	return paciente
}

func DeletarPacientePorID(id int) {
	db.DB.Delete(&model.Paciente{}, id)
}
