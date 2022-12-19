package controller

import (
	"net/http"
	"strconv"

	"github.com/ALTbruno/consultorio-golang/internal/dto"
	"github.com/ALTbruno/consultorio-golang/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary Registra um dentista
// @Description Registra um dentista
// @Accept  json
// @Produce  json
// @Param request body dto.Dentista true "Todos os campos são obrigatórios"
// @Success 201 {object} dto.DentistaResponse
// @Failure 400 {object} dto.Resposta
// @Router /dentistas [post]
func DentistaPOST(c *gin.Context) {
	var dentista dto.Dentista
	if err := c.ShouldBindJSON(&dentista); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, s := service.CadastrarDentista(dentista)
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(201, result)
}

// @Summary Busca um dentista
// @Description Busca um dentista por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Dentista"
// @Success 200 {object} dto.DentistaResponse
// @Failure 400 {object} dto.Resposta
// @Router /dentistas/{id} [get]
func DentistaGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dentista, res := service.BuscarDentistaPorID(uint(id))
	if len(res) > 0 {
		c.JSON(400, gin.H{
			"mensagem": res,
		})
		return
	}
	c.JSON(200, dentista)
}

// @Summary Atualiza um dentista
// @Description Atualiza um dentista por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Dentista"
// @Param request body dto.Dentista true "Todos os campos são obrigatórios"
// @Success 200 {object} dto.DentistaResponse
// @Failure 400 {object} dto.Resposta
// @Router /dentistas/{id} [put]
func DentistaPUT(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dentista dto.Dentista
	if err := c.ShouldBindJSON(&dentista); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, s := service.AtualizarDentistaPorID(dentista, uint(id))
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, d)
}

// @Summary Atualiza um dentista
// @Description Atualiza um dentista por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Dentista"
// @Param request body dto.Dentista true "Não há campos obrigatórios"
// @Success 200 {object} dto.DentistaResponse
// @Failure 400 {object} dto.Resposta
// @Router /dentistas/{id} [patch]
func DentistaPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dentista dto.Dentista
	if err := c.ShouldBindJSON(&dentista); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, s := service.AtualizarDentistaParcial(dentista, uint(id))
	if s != "" {
		c.JSON(400, gin.H{
			"mensagem": s,
		})
		return
	}
	c.JSON(200, d)
}

// @Summary Apaga um dentista
// @Description Apaga um dentista por ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID Dentista"
// @Success 200 {object} dto.Resposta
// @Failure 400 {object} dto.Resposta
// @Router /dentistas/{id} [delete]
func DentistaDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, result := service.DeletarDentistaPorID(uint(id))
	c.JSON(code, gin.H{
		"mensagem": result,
	})
}
