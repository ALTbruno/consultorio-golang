package controller

import (
	"net/http"
	"strconv"

	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/service"
	"github.com/gin-gonic/gin"
)

func DentistaPOST(c *gin.Context) {
	var dentista model.Dentista
	if err := c.ShouldBindJSON(&dentista); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := service.CadastrarDentista(dentista)
	c.JSON(201, result)
}

func DentistaGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := service.BuscarDentistaPorID(id)
	c.JSON(200, result)
}

func DentistaPUT(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "pong",
	})
}

func DentistaPATCH(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "pong",
	})
}

func DentistaDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.DeletarDentistaPorID(id)
	c.Writer.WriteHeader(http.StatusNoContent)
}
