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
	result, s := service.CadastrarDentista(dentista)
	if len(s) > 0 {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dentista model.Dentista
	if err := c.ShouldBindJSON(&dentista); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, s := service.AtualizarDentistaPorID(dentista, id)
	if len(s) > 0 {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, d)
}

func DentistaPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dentista model.Dentista
	if err := c.ShouldBindJSON(&dentista); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, s := service.AtualizarDentistaParcial(dentista, id)
	if len(s) > 0 {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, d)
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
