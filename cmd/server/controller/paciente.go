package controller

import (
	"net/http"
	"strconv"

	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/service"
	"github.com/gin-gonic/gin"
)

func PacientePOST(c *gin.Context) {
	var paciente model.Paciente
	if err := c.ShouldBindJSON(&paciente); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := service.CadastrarPaciente(paciente)
	c.JSON(201, result)
}

func PacienteGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := service.BuscarPacientePorID(id)
	c.JSON(200, result)
}

func PacientePUT(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "pong",
	})
}

func PacientePATCH(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "pong",
	})
}

func PacienteDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.DeletarPacientePorID(id)
	c.Writer.WriteHeader(http.StatusNoContent)
}
