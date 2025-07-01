package controllers

import (
	"cha-casa-nova-backend/database"
	"cha-casa-nova-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetItems retorna todos os itens (visão pública)
func GetItems(c *gin.Context) {
	var items []models.Item
	if err := database.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erro ao buscar itens",
			Code:  http.StatusInternalServerError,
		})
		return
	}

	// Converter para visão pública
	publicItems := make([]models.ItemPublico, len(items))
	for i, item := range items {
		publicItems[i] = item.ToPublic()
	}

	c.JSON(http.StatusOK, publicItems)
}

// GetItem retorna um item específico (visão pública)
func GetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Item não encontrado",
			Code:  http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, item.ToPublic())
}

// CreateItem cria um novo item (admin)
func CreateItem(c *gin.Context) {
	var request models.CreateItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Dados inválidos",
			Code:  http.StatusBadRequest,
		})
		return
	}

	item := models.Item{
		Nome:      request.Nome,
		Descricao: request.Descricao,
		Categoria: request.Categoria,
		Preco:     request.Preco,
		ImagemURL: request.ImagemURL,
		LinkURL:   request.LinkURL,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erro ao criar item",
			Code:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, models.AdminResponse{
		Success: true,
		Message: "Item criado com sucesso",
		Data:    item,
	})
}

// UpdateItem atualiza um item existente (admin)
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Item não encontrado",
			Code:  http.StatusNotFound,
		})
		return
	}

	var request models.CreateItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
		return
	}

	item.Nome = request.Nome
	item.Descricao = request.Descricao
	item.Categoria = request.Categoria
	item.Preco = request.Preco
	item.ImagemURL = request.ImagemURL
	item.LinkURL = request.LinkURL

	if err := database.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erro ao atualizar item",
			Code:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.AdminResponse{
		Success: true,
		Message: "Item atualizado com sucesso",
		Data:    item,
	})
}

// DeleteItem remove um item (admin)
func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Item{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erro ao deletar item",
			Code:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.AdminResponse{
		Success: true,
		Message: "Item deletado com sucesso",
	})
}

// ResgateItem marca um item como resgatado (público)
func ResgateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Item não encontrado",
			Code:  http.StatusNotFound,
		})
		return
	}

	if item.Resgatado {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Item já foi resgatado",
			Code:  http.StatusBadRequest,
		})
		return
	}

	var request models.ResgatarItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Nome é obrigatório",
			Code:  http.StatusBadRequest,
		})
		return
	}

	now := time.Now()
	item.Resgatado = true
	item.ResgatadoPor = request.Nome
	item.ResgatadoEm = now

	if err := database.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erro ao resgatar item",
			Code:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item resgatado com sucesso!",
		"item":    item.ToPublic(),
	})
}

// CancelaResgate desmarca um item como resgatado (público)
func CancelaResgate(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Item não encontrado",
			Code:  http.StatusNotFound,
		})
		return
	}

	if !item.Resgatado {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Item não está resgatado",
			Code:  http.StatusBadRequest,
		})
		return
	}

	item.Resgatado = false
	item.ResgatadoPor = ""
	item.ResgatadoEm = time.Time{} // Zero value para time.Time

	if err := database.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erro ao cancelar resgate",
			Code:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Resgate cancelado com sucesso!",
		"item":    item.ToPublic(),
	})
}

// GetAdminItems retorna todos os itens (visão administrativa)
func GetAdminItems(c *gin.Context) {
	var items []models.Item
	if err := database.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erro ao buscar itens",
			Code:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, items)
}
