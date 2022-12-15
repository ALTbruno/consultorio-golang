package main

import (
	"github.com/ALTbruno/consultorio-golang/config"
	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/ALTbruno/consultorio-golang/internal/model"
)

func init() {
	config.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(&model.Dentista{})
	db.DB.AutoMigrate(&model.Paciente{})
	db.DB.AutoMigrate(&model.Consulta{})
}
