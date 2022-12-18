package service

import (
	"fmt"
	"time"

	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/repository"
	"github.com/go-playground/validator/v10"
)

func CadastrarPaciente(paciente model.Paciente) (model.Paciente, string) {
	validate := validator.New()
	err := validate.Struct(paciente)
	if err != nil {
		return model.Paciente{}, err.Error()
	}
	hoje := time.Now().Format("02-01-2006")
	paciente.DataDeCadastro = hoje
	p := repository.CadastrarPaciente(paciente)
	return p, ""
}

func BuscarPacientePorID(id int) (model.Paciente, string) {
	if !repository.ExistePacientePorId(id) {
		return model.Paciente{}, fmt.Sprintf("Paciente não encontrado %d", id)
	}
	paciente := repository.BuscarPacientePorID(id)
	return paciente, ""
}

func BuscarPacientePorRG(rg string) (model.Paciente, string) {
	paciente := repository.BuscarPacientePorRG(rg)
	if paciente == (model.Paciente{}) {
		return model.Paciente{}, fmt.Sprintf("Paciente não encontrado com o RG %s", rg)
	}
	return paciente, ""
}

func AtualizarPacientePorID(paciente model.Paciente, id int) (model.Paciente, string) {
	validate := validator.New()
	err := validate.Struct(paciente)
	if err != nil {
		return model.Paciente{}, err.Error()
	}
	pacienteSalvo, s := BuscarPacientePorID(id)
	if s != "" {
		return model.Paciente{}, s
	}
	pacienteSalvo.Nome = paciente.Nome
	pacienteSalvo.Sobrenome = paciente.Sobrenome
	pacienteSalvo.RG = paciente.RG
	p := repository.AtualizarPaciente(pacienteSalvo)
	return p, ""
}

func AtualizarPacienteParcial(paciente model.Paciente, id int) (model.Paciente, string) {
	pacienteSalvo, s := BuscarPacientePorID(id)
	if s != "" {
		return model.Paciente{}, s
	}
	colunas := make(map[string]interface{})
	if paciente.Nome != "" {
		colunas["nome"] = paciente.Nome
	}
	if paciente.Sobrenome != "" {
		colunas["sobrenome"] = paciente.Sobrenome
	}
	if paciente.RG != "" {
		colunas["rg"] = paciente.RG
	}

	p := repository.AtualizarPacienteParcial(pacienteSalvo, colunas)
	return p, ""
}

func DeletarPacientePorID(id int) (int, string) {
	paciente, res := BuscarPacientePorID(id)
	if res != "" {
		return 400, res
	}
	result := repository.DeletarPaciente(paciente)
	if result {
		return 200, fmt.Sprintf("Paciente %d deletado com sucesso", id)
	} else {
		return 500, fmt.Sprintf("Erro ao deletar o paciente %d", id)
	}
}
