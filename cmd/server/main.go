package main

import (
	"github.com/ALTbruno/consultorio-golang/cmd/server/controller"
	"github.com/ALTbruno/consultorio-golang/config"
	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {

	app := gin.Default()

	dentistas := app.Group("/dentistas")
	{
		dentistas.POST("", controller.DentistaPOST)
		dentistas.GET("/:id", controller.DentistaGET)
		dentistas.PUT("/:id", controller.DentistaPUT)
		dentistas.PATCH("/:id", controller.DentistaPATCH)
		dentistas.DELETE("/:id", controller.DentistaDELETE)
	}

	pacientes := app.Group("/pacientes")
	{
		pacientes.POST("", controller.PacientePOST)
		pacientes.GET("/:id", controller.PacienteGET)
		pacientes.PUT("/:id", controller.PacientePUT)
		pacientes.PATCH("/:id", controller.PacientePATCH)
		pacientes.DELETE("/:id", controller.PacienteDELETE)
	}

	consultas := app.Group("/consultas")
	{
		consultas.POST("", controller.ConsultaPOST)
		consultas.GET("/:id", controller.ConsultaGET)
		consultas.PUT("/:id", controller.ConsultaPUT)
		consultas.PATCH("/:id", controller.ConsultaPATCH)
		consultas.DELETE("/:id", controller.ConsultaDELETE)
		consultas.POST("/marcar", controller.ConsultaPorMatriculaRGPOST)
	}

	app.Run()
}