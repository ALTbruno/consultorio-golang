package repository

import (
	"log"

	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/ALTbruno/consultorio-golang/internal/model"
	"gorm.io/gorm"
)

func CadastrarPaciente(paciente model.Paciente) model.Paciente {
	db.DB.Create(&paciente)
	return paciente
}

func BuscarPacientePorID(id uint) model.Paciente {
	var paciente model.Paciente
	db.DB.First(&paciente, id)
	return paciente
}

func BuscarPacientePorRG(rg string) model.Paciente {
	var paciente model.Paciente
	db.DB.Where("rg = ?", rg).First(&paciente)
	return paciente
}

func ExistePacientePorId(id uint) bool {
	var paciente model.Paciente
	err := db.DB.First(&paciente, id).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}

func DeletarPacientePorId(id uint) bool {
	result := db.DB.Delete(&model.Paciente{}, id)
	return result.RowsAffected > 0
}

func AtualizarPacientePorId(id uint, colunas map[string]interface{}) model.Paciente {
	paciente := BuscarPacientePorID(id)
	db.DB.Model(&paciente).Where("id = ?", id).UpdateColumns(colunas)
	return paciente
}
