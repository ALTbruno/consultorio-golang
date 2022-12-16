package service

import (
	"fmt"

	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/repository"
)

func CadastrarDentista(dentista model.Dentista) model.Dentista {
	d := repository.CadastrarDentista(dentista)
	return d
}

func BuscarDentistaPorID(id int) (model.Dentista, string) {
	if !repository.ExisteDentistaPorId(id) {
		return model.Dentista{}, fmt.Sprintf("Dentista não encontrado %d", id)
	}
	dentista := repository.BuscarDentistaPorID(id)
	return dentista, ""
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
