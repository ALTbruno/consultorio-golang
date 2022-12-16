package repository

import (
	"log"

	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/ALTbruno/consultorio-golang/internal/model"
	"gorm.io/gorm"
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

func ExisteDentistaPorId(id int) bool {
	var dentista model.Dentista
	err := db.DB.First(&dentista, id).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}

func AtualizarDentistaPorID(id int) model.Dentista {
	dentista := BuscarDentistaPorID(id)
	db.DB.First(&dentista, id)
	return dentista
}

func DeletarDentista(dentista model.Dentista) bool {
	result := db.DB.Delete(&dentista)
	return result.RowsAffected > 0
}
