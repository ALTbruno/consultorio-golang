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

func BuscarDentistaPorID(id int) model.Dentista {
	dentista := repository.BuscarDentistaPorID(id)
	return dentista
}

func DeletarDentistaPorID(id int) (int, string) {
	if !repository.ExisteDentistaPorId(id) {
		return 400, fmt.Sprintf("Dentista %d não encontrado", id)
	}
	dentista := BuscarDentistaPorID(id)
	result := repository.DeletarDentista(dentista)
	if result {
		return 200, fmt.Sprintf("Dentista %d deletado com sucesso", id)
	} else {
		return 500, fmt.Sprintf("Erro ao deletar o dentista %d", id)
	}
}
