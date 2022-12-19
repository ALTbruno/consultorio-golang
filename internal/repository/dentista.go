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

func BuscarDentistaPorID(id uint) model.Dentista {
	var dentista model.Dentista
	db.DB.First(&dentista, id)
	return dentista
}

func BuscarDentistaPorMatricula(matricula string) model.Dentista {
	var dentista model.Dentista
	db.DB.Where("matricula = ?", matricula).First(&dentista)
	return dentista
}

func ExisteDentistaPorId(id uint) bool {
	var dentista model.Dentista
	err := db.DB.First(&dentista, id).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}

func DeletarDentistaPorId(id uint) bool {
	result := db.DB.Delete(&model.Dentista{}, id)
	return result.RowsAffected > 0
}

func AtualizarDentistaPorId(id uint, colunas map[string]interface{}) model.Dentista {
	dentista := BuscarDentistaPorID(id)
	db.DB.Model(&dentista).Where("id = ?", id).UpdateColumns(colunas)
	return dentista
}
