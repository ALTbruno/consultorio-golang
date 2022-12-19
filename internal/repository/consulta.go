package repository

import (
	"log"

	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/ALTbruno/consultorio-golang/internal/model"
	"gorm.io/gorm"
)

func CadastrarConsulta(consulta model.Consulta) model.Consulta {
	db.DB.Create(&consulta)
	return consulta
}

func ExisteConsultaPorId(id uint) bool {
	var consulta model.Consulta
	err := db.DB.First(&consulta, id).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}

func BuscarConsultaPorID(id uint) model.Consulta {
	var consulta model.Consulta
	db.DB.First(&consulta, id)
	return consulta
}

func BuscarConsultasPorIdPaciente(id uint) []model.Consulta {
	var consultas []model.Consulta
	db.DB.Where("paciente_id = ?", id).Find(&consultas)
	return consultas
}

func DeletarConsultaPorId(id uint) bool {
	result := db.DB.Delete(&model.Consulta{}, id)
	return result.RowsAffected > 0
}

func AtualizarConsultaPorId(id uint, colunas map[string]interface{}) model.Consulta {
	consulta := BuscarConsultaPorID(id)
	db.DB.Model(&consulta).Where("id = ?", id).UpdateColumns(colunas)
	return consulta
}
