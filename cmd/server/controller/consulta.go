package controller

import (
	"net/http"
	"strconv"

	"github.com/ALTbruno/consultorio-golang/internal/dto"
	"github.com/ALTbruno/consultorio-golang/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary Registra uma consulta
// @Description Registra uma consulta
// @Accept  json
// @Produce  json
// @Param request body dto.Consulta true "Todos os campos são obrigatórios"
// @Success 201 {object} dto.ConsultaResponse
// @Failure 400 {object} dto.Resposta
// @Router /consultas [post]
func ConsultaPOST(c *gin.Context) {
	var consulta dto.Consulta
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

// @Summary Registra uma consulta
// @Description Registra uma consulta com matricula_dentista e rg_paciente
// @Accept  json
// @Produce  json
// @Param request body dto.ConsultaMatriculaRG true "Todos os campos são obrigatórios"
// @Success 201 {object} dto.ConsultaResponse
// @Failure 400 {object} dto.Resposta
// @Router /consultas/marcar [post]
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

// @Summary Busca uma consulta
// @Description Busca uma consulta por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Consulta"
// @Success 200 {object} dto.ConsultaResponse
// @Failure 400 {object} dto.Resposta
// @Router /consultas/{id} [get]
func ConsultaGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, s := service.BuscarConsultaPorID(uint(id))
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

// @Summary Busca consultas
// @Description Busca consultas por RG do paciente
// @Accept  json
// @Produce  json
// @Param   rg_paciente     query    string     true        "RG do Paciente"
// @Success 200 {object} []dto.ConsultaResponse
// @Failure 400 {object} dto.Resposta
// @Router /consultas [get]
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

// @Summary Atualiza uma consulta
// @Description Atualiza uma consulta por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Consulta"
// @Param request body dto.Consulta true "Todos os campos são obrigatórios"
// @Success 200 {object} dto.ConsultaResponse
// @Failure 400 {object} dto.Resposta
// @Router /consultas/{id} [put]
func ConsultaPUT(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var consulta dto.Consulta
	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.AtualizarConsultaPorID(consulta, uint(id))
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

// @Summary Atualiza uma consulta
// @Description Atualiza uma consulta por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Consulta"
// @Param request body dto.Consulta true "Não há campos obrigatórios"
// @Success 200 {object} dto.ConsultaResponse
// @Failure 400 {object} dto.Resposta
// @Router /consultas/{id} [patch]
func ConsultaPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var consulta dto.Consulta
	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.AtualizarConsultaParcial(consulta, uint(id))
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

// @Summary Apaga uma consulta
// @Description Apaga uma consulta por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Consulta"
// @Success 200 {object} dto.Resposta
// @Failure 400 {object} dto.Resposta
// @Router /consultas/{id} [delete]
func ConsultaDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, result := service.DeletarConsultaPorID(uint(id))
	c.JSON(code, gin.H{
		"mensagem": result,
	})
}
