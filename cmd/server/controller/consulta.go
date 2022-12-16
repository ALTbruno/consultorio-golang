package controller

import (
	"net/http"
	"strconv"

	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/service"
	"github.com/gin-gonic/gin"
)

func ConsultaPOST(c *gin.Context) {
	var consulta model.Consulta
	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := service.CadastrarConsulta(consulta)
	c.JSON(201, result)
}

func ConsultaGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := service.BuscarConsultaPorID(id)
	c.JSON(200, result)
}

func ConsultaPUT(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "pong",
	})
}

func ConsultaPATCH(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "pong",
	})
}

func ConsultaDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.DeletarConsultaPorID(id)
	c.Writer.WriteHeader(http.StatusNoContent)
}
