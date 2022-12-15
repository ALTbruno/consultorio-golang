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

	app.Run()
}