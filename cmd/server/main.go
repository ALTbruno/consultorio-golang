package main

import (
	"fmt"
	"os"

	"github.com/ALTbruno/consultorio-golang/cmd/server/controller"
	"github.com/ALTbruno/consultorio-golang/cmd/server/docs"
	"github.com/ALTbruno/consultorio-golang/config"
	"github.com/ALTbruno/consultorio-golang/db"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func init() {
	config.LoadEnvVariables()
	db.ConnectToDB()
}

// @title Cosultório Odontológico
// @version 1.0
// @description API responsável pelo gerenciamento de Dentistas, Pacientes e Consultas
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	app := gin.Default()

	apiHost := os.Getenv("HOST")
	apiPort := os.Getenv("PORT")
	host := fmt.Sprintf("%s:%s", apiHost, apiPort)
	docs.SwaggerInfo.Host = host

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		consultas.GET("", controller.BuscarConsultasPorRGPaciente)
		consultas.GET("/:id", controller.ConsultaGET)
		consultas.PUT("/:id", controller.ConsultaPUT)
		consultas.PATCH("/:id", controller.ConsultaPATCH)
		consultas.DELETE("/:id", controller.ConsultaDELETE)
		consultas.POST("/marcar", controller.ConsultaPorMatriculaRGPOST)
	}

	app.Run()
}