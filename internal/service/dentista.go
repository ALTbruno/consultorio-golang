package service

import (
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

func DeletarDentistaPorID(id int) {
	repository.DeletarDentistaPorID(id)
}
