package controllers

import (
	"net/http"
	"time"

	"cha-casa-nova-backend/database"
	"cha-casa-nova-backend/models"

	"github.com/gin-gonic/gin"
)

// GetConfirmacoes retorna todas as confirmações
func GetConfirmacoes(c *gin.Context) {
	var confirmacoes []models.Confirmacao
	if err := database.DB.Find(&confirmacoes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar confirmações"})
		return
	}
	c.JSON(http.StatusOK, confirmacoes)
}

// CreateConfirmacao cria uma nova confirmação
func CreateConfirmacao(c *gin.Context) {
	var confirmacao models.Confirmacao
	if err := c.ShouldBindJSON(&confirmacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	confirmacao.CriadaEm = time.Now()

	if err := database.DB.Create(&confirmacao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar confirmação"})
		return
	}

	c.JSON(http.StatusCreated, confirmacao)
}

// UpdateConfirmacao - Atualizar confirmação
func UpdateConfirmacao(c *gin.Context) {
	id := c.Param("id")
	var confirmacao models.Confirmacao

	if err := database.DB.First(&confirmacao, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Confirmação não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&confirmacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&confirmacao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar confirmação"})
		return
	}

	c.JSON(http.StatusOK, confirmacao)
}

// DeleteConfirmacao - Deletar confirmação
func DeleteConfirmacao(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Confirmacao{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar confirmação"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Confirmação deletada com sucesso"})
}
