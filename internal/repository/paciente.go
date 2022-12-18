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

func BuscarPacientePorID(id int) model.Paciente {
	var paciente model.Paciente
	db.DB.First(&paciente, id)
	return paciente
}

func BuscarPacientePorRG(rg string) model.Paciente {
	var paciente model.Paciente
	db.DB.Where("rg = ?", rg).First(&paciente)
	return paciente
}

func ExistePacientePorId(id int) bool {
	var paciente model.Paciente
	err := db.DB.First(&paciente, id).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}

func AtualizarPaciente(paciente model.Paciente) model.Paciente {
	db.DB.Save(&paciente)
	return paciente
}

func DeletarPaciente(paciente model.Paciente) bool {
	result:= db.DB.Delete(&paciente)
	return result.RowsAffected > 0
}

func AtualizarPacienteParcial(paciente model.Paciente, colunas map[string]interface{}) model.Paciente {
	db.DB.Model(&paciente).UpdateColumns(colunas)
	return paciente
}
