package service

import (
	"fmt"
	"time"

	"github.com/ALTbruno/consultorio-golang/internal/dto"
	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/repository"
	"github.com/go-playground/validator/v10"
)

func CadastrarPaciente(paciente dto.Paciente) (dto.PacienteResponse, string) {
	validate := validator.New()
	err := validate.Struct(paciente)
	if err != nil {
		return dto.PacienteResponse{}, err.Error()
	}
	hoje := time.Now().Format("02-01-2006")
	p := repository.CadastrarPaciente(model.Paciente{
		Nome: paciente.Nome,
		Sobrenome: paciente.Sobrenome,
		RG: paciente.Sobrenome,
		DataDeCadastro: hoje,
	})
	pacienteResponse := ConverterPacienteModelParaResponse(p)
	return pacienteResponse, ""
}

func BuscarPacientePorID(id uint) (dto.PacienteResponse, string) {
	if !repository.ExistePacientePorId(id) {
		return dto.PacienteResponse{}, fmt.Sprintf("Paciente não encontrado %d", id)
	}
	paciente := repository.BuscarPacientePorID(id)
	pacienteResponse := ConverterPacienteModelParaResponse(paciente)
	return pacienteResponse, ""
}

func BuscarPacientePorRG(rg string) (dto.PacienteResponse, string) {
	paciente := repository.BuscarPacientePorRG(rg)
	if paciente == (model.Paciente{}) {
		return dto.PacienteResponse{}, fmt.Sprintf("Paciente não encontrado com o RG %s", rg)
	}
	pacienteResponse := ConverterPacienteModelParaResponse(paciente)
	return pacienteResponse, ""
}

func AtualizarPacientePorID(paciente dto.Paciente, id uint) (dto.PacienteResponse, string) {
	validate := validator.New()
	err := validate.Struct(paciente)
	if err != nil {
		return dto.PacienteResponse{}, err.Error()
	}
	_, s := BuscarPacientePorID(id)
	if s != "" {
		return dto.PacienteResponse{}, s
	}
	colunas := make(map[string]interface{})
	colunas["nome"] = paciente.Nome
	colunas["sobrenome"] = paciente.Sobrenome
	colunas["rg"] = paciente.RG
	p := repository.AtualizarPacientePorId(id, colunas)
	pacienteResponse := ConverterPacienteModelParaResponse(p)
	return pacienteResponse, ""
}

func AtualizarPacienteParcial(paciente dto.Paciente, id uint) (dto.PacienteResponse, string) {
	_, s := BuscarPacientePorID(id)
	if s != "" {
		return dto.PacienteResponse{}, s
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

	p := repository.AtualizarPacientePorId(id, colunas)
	pacienteResponse := ConverterPacienteModelParaResponse(p)
	return pacienteResponse, ""
}

func DeletarPacientePorID(id uint) (int, string) {
	_, res := BuscarPacientePorID(id)
	if res != "" {
		return 400, res
	}
	result := repository.DeletarPacientePorId(id)
	if result {
		return 200, fmt.Sprintf("Paciente %d deletado com sucesso", id)
	} else {
		return 500, fmt.Sprintf("Erro ao deletar o paciente %d", id)
	}
}

func ConverterPacienteModelParaResponse(paciente model.Paciente) dto.PacienteResponse {
	return dto.PacienteResponse{
		ID: paciente.ID,
		Nome: paciente.Nome,
		Sobrenome: paciente.Sobrenome,
		RG: paciente.RG,
		DataDeCadastro: paciente.DataDeCadastro,
	}
}
