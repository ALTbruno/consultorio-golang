package service

import (
	"fmt"

	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/repository"
	"github.com/go-playground/validator/v10"
)

func CadastrarDentista(dentista model.Dentista) (model.Dentista, string) {
	validate := validator.New()
	err := validate.Struct(dentista)
	if err != nil {
		return model.Dentista{}, err.Error()
	}
	d := repository.CadastrarDentista(dentista)
	return d, ""
}

func BuscarDentistaPorID(id int) (model.Dentista, string) {
	if !repository.ExisteDentistaPorId(id) {
		return model.Dentista{}, fmt.Sprintf("Dentista não encontrado %d", id)
	}
	dentista := repository.BuscarDentistaPorID(id)
	return dentista, ""
}

func AtualizarDentistaPorID(dentista model.Dentista, id int) (model.Dentista, string) {
	validate := validator.New()
	err := validate.Struct(dentista)
	if err != nil {
		return model.Dentista{}, err.Error()
	}
	dentistaSalvo, s := BuscarDentistaPorID(id)
	if len(s) > 0 {
		return model.Dentista{}, s
	}
	dentistaSalvo.Nome = dentista.Nome
	dentistaSalvo.Sobrenome = dentista.Sobrenome
	dentistaSalvo.Matricula = dentista.Matricula
	d := repository.AtualizarDentista(dentistaSalvo)
	return d, ""
}

func AtualizarDentistaParcial(dentista model.Dentista, id int) (model.Dentista, string) {
	dentistaSalvo, s := BuscarDentistaPorID(id)
	if len(s) > 0 {
		return model.Dentista{}, s
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
	
	r := repository.AtualizarDentistaParcial(dentistaSalvo, colunas)
	return r, ""
}

func DeletarDentistaPorID(id int) (int, string) {
	dentista, res := BuscarDentistaPorID(id)
	if len(res) > 0 {
		return 400, res
	}
	result := repository.DeletarDentista(dentista)
	if result {
		return 200, fmt.Sprintf("Dentista %d deletado com sucesso", id)
	} else {
		return 500, fmt.Sprintf("Erro ao deletar o dentista %d", id)
	}
}
