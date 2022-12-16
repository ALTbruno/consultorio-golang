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

	dentista, res := service.BuscarDentistaPorID(id)
	if len(res) > 0 {
		c.JSON(400, gin.H{
			"mensagem": res,
		})
		return
	}
	c.JSON(200, dentista)
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
	code, result := service.DeletarDentistaPorID(id)
	c.JSON(code, gin.H{
		"mensagem": result,
	})
}
