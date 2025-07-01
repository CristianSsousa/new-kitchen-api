package controllers

import (
	"net/http"
	"strconv"
	"time"

	"cha-casa-nova-backend/database"
	"cha-casa-nova-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetMensagens retorna todas as mensagens
func GetMensagens(c *gin.Context) {
	var mensagens []models.Mensagem
	if err := database.DB.Find(&mensagens).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar mensagens"})
		return
	}
	c.JSON(http.StatusOK, mensagens)
}

// CreateMensagem cria uma nova mensagem
func CreateMensagem(c *gin.Context) {
	var mensagem models.Mensagem
	if err := c.ShouldBindJSON(&mensagem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	mensagem.CriadaEm = time.Now()
	mensagem.Aprovada = false

	if err := database.DB.Create(&mensagem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar mensagem"})
		return
	}

	c.JSON(http.StatusCreated, mensagem)
}

// AprovarMensagem aprova uma mensagem
func AprovarMensagem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var mensagem models.Mensagem
	if err := database.DB.First(&mensagem, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mensagem não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar mensagem"})
		return
	}

	mensagem.Aprovada = true
	mensagem.AprovadaEm = time.Now()

	if err := database.DB.Save(&mensagem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao aprovar mensagem"})
		return
	}

	c.JSON(http.StatusOK, mensagem)
}

// DeleteMensagem exclui uma mensagem
func DeleteMensagem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var mensagem models.Mensagem
	if err := database.DB.First(&mensagem, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mensagem não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar mensagem"})
		return
	}

	if err := database.DB.Delete(&mensagem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir mensagem"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetMensagensAprovadas retorna apenas as mensagens aprovadas
func GetMensagensAprovadas(c *gin.Context) {
	var mensagens []models.Mensagem
	if err := database.DB.Where("aprovada = ?", true).Find(&mensagens).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar mensagens"})
		return
	}
	c.JSON(http.StatusOK, mensagens)
}
