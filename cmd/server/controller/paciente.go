package controller

import (
	"net/http"
	"strconv"

	"github.com/ALTbruno/consultorio-golang/internal/dto"
	"github.com/ALTbruno/consultorio-golang/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary Registra um paciente
// @Description Registra um paciente
// @Accept  json
// @Produce  json
// @Param request body dto.Paciente true "Todos os campos são obrigatórios"
// @Success 201 {object} dto.PacienteResponse
// @Failure 400 {object} dto.Resposta
// @Router /pacientes [post]
func PacientePOST(c *gin.Context) {
	var paciente dto.Paciente
	if err := c.ShouldBindJSON(&paciente); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.CadastrarPaciente(paciente)
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(201, result)
}

// @Summary Busca um paciente
// @Description Busca um paciente por ID
// @Accept  json
// @Produce  json
// @Param id path int true "ID Paciente"
// @Success 200 {object} dto.PacienteResponse
// @Failure 400 {object} dto.Resposta
// @Router /pacientes/{id} [get]
func PacienteGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, s := service.BuscarPacientePorID(uint(id))
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

// @Summary Atualiza um paciente
// @Description Atualiza um paciente por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Paciente"
// @Param request body dto.Paciente true "Todos os campos são obrigatórios"
// @Success 200 {object} dto.PacienteResponse
// @Failure 400 {object} dto.Resposta
// @Router /pacientes/{id} [put]
func PacientePUT(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var paciente dto.Paciente
	if err := c.ShouldBindJSON(&paciente); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.AtualizarPacientePorID(paciente, uint(id))
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

// @Summary Atualiza um paciente
// @Description Atualiza um paciente por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Paciente"
// @Param request body dto.Paciente true "Não há campos obrigatórios"
// @Success 200 {object} dto.PacienteResponse
// @Failure 400 {object} dto.Resposta
// @Router /pacientes/{id} [patch]
func PacientePATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var paciente dto.Paciente
	if err := c.ShouldBindJSON(&paciente); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.AtualizarPacienteParcial(paciente, uint(id))
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, result)
}

// @Summary Apaga um paciente
// @Description Apaga um paciente por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Paciente"
// @Success 200 {object} dto.Resposta
// @Failure 400 {object} dto.Resposta
// @Router /pacientes/{id} [delete]
func PacienteDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, result := service.DeletarPacientePorID(uint(id))
	c.JSON(code, gin.H{
		"mensagem": result,
	})
}
