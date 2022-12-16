package repository

import (
	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/ALTbruno/consultorio-golang/internal/model"
)

func CadastrarConsulta(consulta model.Consulta) model.Consulta {
	db.DB.Create(&consulta)
	return consulta
}

func BuscarConsultaPorID(id int) model.Consulta {
	var consulta model.Consulta
	db.DB.First(&consulta, id)
	return consulta
}

func AtualizarConsultaPorID(id int) model.Consulta {
	consulta := BuscarConsultaPorID(id)
	db.DB.First(&consulta, id)
	return consulta
}

func DeletarConsultaPorID(id int) {
	db.DB.Delete(&model.Consulta{}, id)
}
