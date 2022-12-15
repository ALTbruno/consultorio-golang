package repository

import (
	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/ALTbruno/consultorio-golang/internal/model"
)

func CadastrarDentista(dentista model.Dentista) model.Dentista {
	db.DB.Create(&dentista)
	return dentista
}

func BuscarDentistaPorID(id int) model.Dentista {
	var dentista model.Dentista
	db.DB.First(&dentista, id)
	return dentista
}

func AtualizarDentistaPorID(id int) model.Dentista {
	dentista := BuscarDentistaPorID(id)
	db.DB.First(&dentista, id)
	return dentista
}

func DeletarDentistaPorID(id int) {
	db.DB.Delete(&model.Dentista{}, id)
}
