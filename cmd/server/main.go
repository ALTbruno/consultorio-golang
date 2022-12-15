package main

import (
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
	app.Run()
}