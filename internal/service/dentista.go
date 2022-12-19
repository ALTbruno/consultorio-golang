package service

import (
	"fmt"

	"github.com/ALTbruno/consultorio-golang/internal/dto"
	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/repository"
	"github.com/go-playground/validator/v10"
)

func CadastrarDentista(dentista dto.Dentista) (dto.DentistaResponse, string) {
	validate := validator.New()
	err := validate.Struct(dentista)
	if err != nil {
		return dto.DentistaResponse{}, err.Error()
	}
	d := repository.CadastrarDentista(
		model.Dentista{
			Nome: dentista.Nome,
			Sobrenome: dentista.Sobrenome,
			Matricula: dentista.Matricula,
		},
	)
	dentistaResponse := ConverteDentistaModelParaResponse(d)
	return dentistaResponse, ""
}

func BuscarDentistaPorID(id uint) (dto.DentistaResponse, string) {
	if !repository.ExisteDentistaPorId(id) {
		return dto.DentistaResponse{}, fmt.Sprintf("Dentista não encontrado %d", id)
	}
	dentista := repository.BuscarDentistaPorID(id)
	dentistaResponse := ConverteDentistaModelParaResponse(dentista)
	return dentistaResponse, ""
}

func BuscarDentistaPorMatricula(matricula string) (dto.DentistaResponse, string) {
	dentista := repository.BuscarDentistaPorMatricula(matricula)
	if dentista == (model.Dentista{}) {
		return dto.DentistaResponse{}, fmt.Sprintf("Dentista não encontrado com a Matrícula %s", matricula)
	}
	dentistaResponse := ConverteDentistaModelParaResponse(dentista)
	return dentistaResponse, ""
}

func AtualizarDentistaPorID(dentista dto.Dentista, id uint) (dto.DentistaResponse, string) {
	validate := validator.New()
	err := validate.Struct(dentista)
	if err != nil {
		return dto.DentistaResponse{}, err.Error()
	}
	_, s := BuscarDentistaPorID(id)
	if s != "" {
		return dto.DentistaResponse{}, s
	}
	colunas := make(map[string]interface{})
	colunas["nome"] = dentista.Nome
	colunas["sobrenome"] = dentista.Sobrenome
	colunas["matricula"] = dentista.Matricula
	d := repository.AtualizarDentistaPorId(id, colunas)
	dentistaResponse := ConverteDentistaModelParaResponse(d)
	return dentistaResponse, ""
}

func AtualizarDentistaParcial(dentista dto.Dentista, id uint) (dto.DentistaResponse, string) {
	_, s := BuscarDentistaPorID(id)
	if s != "" {
		return dto.DentistaResponse{}, s
	}
	colunas := make(map[string]interface{})
	if dentista.Nome != "" {
		colunas["nome"] = dentista.Nome
	}
	if dentista.Sobrenome != "" {
		colunas["sobrenome"] = dentista.Sobrenome
	}
	if dentista.Matricula != "" {
		colunas["matricula"] = dentista.Matricula
	}
	
	d := repository.AtualizarDentistaPorId(id, colunas)
	dentistaResponse := ConverteDentistaModelParaResponse(d)
	return dentistaResponse, ""
}

func DeletarDentistaPorID(id uint) (int, string) {
	_, res := BuscarDentistaPorID(id)
	if res != "" {
		return 400, res
	}
	result := repository.DeletarDentistaPorId(id)
	if result {
		return 200, fmt.Sprintf("Dentista %d deletado com sucesso", id)
	} else {
		return 500, fmt.Sprintf("Erro ao deletar o dentista %d", id)
	}
}

func ConverteDentistaModelParaResponse(dentista model.Dentista) dto.DentistaResponse {
	return dto.DentistaResponse{
		ID: dentista.ID,
		Nome: dentista.Nome,
		Sobrenome: dentista.Sobrenome,
		Matricula: dentista.Matricula,
	}
}
