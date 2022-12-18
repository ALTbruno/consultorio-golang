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

func ExisteConsultaPorId(id int) bool {
	var consulta model.Consulta
	err := db.DB.First(&consulta, id).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}

func BuscarConsultaPorID(id int) model.Consulta {
	var consulta model.Consulta
	db.DB.First(&consulta, id)
	return consulta
}

func AtualizarConsulta(consulta model.Consulta) model.Consulta {
	db.DB.Save(&consulta)
	return consulta
}

func DeletarConsulta(consulta model.Consulta) bool {
	result := db.DB.Delete(&consulta)
	return result.RowsAffected > 0
}

func AtualizarConsultaParcial(consulta model.Consulta, colunas map[string]interface{}) model.Consulta {
	db.DB.Model(&consulta).UpdateColumns(colunas)
	return consulta
}
