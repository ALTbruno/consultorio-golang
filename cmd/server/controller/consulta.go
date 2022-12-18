package controller

import (
	"net/http"
	"strconv"

	"github.com/ALTbruno/consultorio-golang/internal/dto"
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
	result, s := service.CadastrarConsulta(consulta)
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(201, result)
}

func ConsultaPorMatriculaRGPOST(c *gin.Context) {
	var consulta dto.ConsultaMatriculaRG
	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.CadastrarConsultaPorMatriculaRG(consulta)
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(201, result)
}

func ConsultaGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, s := service.BuscarConsultaPorID(id)
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

func BuscarConsultasPorRGPaciente(c *gin.Context) {
	rg := c.Query("rg_paciente")
	result, s := service.BuscarConsultasPorRGPaciente(rg)
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)

}

func ConsultaPUT(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var consulta model.Consulta
	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.AtualizarConsultaPorID(consulta, id)
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

func ConsultaPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var consulta model.Consulta
	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.AtualizarConsultaParcial(consulta, id)
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

func ConsultaDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, result := service.DeletarConsultaPorID(id)
	c.JSON(code, gin.H{
		"mensagem": result,
	})
}
