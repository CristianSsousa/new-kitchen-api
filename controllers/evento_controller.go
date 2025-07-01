package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"cha-casa-nova-backend/models"

	"github.com/gin-gonic/gin"
)

var eventoFilePath string

func init() {
	// Obtém o diretório do projeto
	projectDir, err := os.Getwd()
	if err != nil {
		log.Printf("Erro ao obter diretório atual: %v", err)
		projectDir = "."
	}

	// Define o caminho do arquivo
	eventoFilePath = filepath.Join(projectDir, "data", "evento.json")
	log.Printf("Caminho do arquivo de evento: %s", eventoFilePath)

	// Garante que o diretório data existe
	if err := os.MkdirAll(filepath.Dir(eventoFilePath), os.ModePerm); err != nil {
		log.Printf("Erro ao criar diretório data: %v", err)
	}
}

// GetEvento retorna as informações do evento
func GetEvento(c *gin.Context) {
	log.Printf("Buscando informações do evento em: %s", eventoFilePath)

	// Verifica se o arquivo existe
	if _, err := os.Stat(eventoFilePath); os.IsNotExist(err) {
		log.Printf("Arquivo não encontrado, retornando valores padrão")
		// Se não existir, retorna valores padrão
		c.JSON(http.StatusOK, models.Evento{
			Data:      "",
			Horario:   "",
			Local:     "",
			UpdatedAt: time.Now(),
		})
		return
	}

	// Lê o arquivo
	fileContent, err := os.ReadFile(eventoFilePath)
	if err != nil {
		log.Printf("Erro ao ler arquivo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao ler informações do evento: %v", err)})
		return
	}

	var evento models.Evento
	if err := json.Unmarshal(fileContent, &evento); err != nil {
		log.Printf("Erro ao fazer unmarshal do JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao processar informações do evento: %v", err)})
		return
	}

	c.JSON(http.StatusOK, evento)
}

// UpdateEvento atualiza as informações do evento
func UpdateEvento(c *gin.Context) {
	var request models.UpdateEventoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Erro ao fazer bind do JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Dados inválidos: %v", err)})
		return
	}

	// Atualiza o evento
	evento := models.Evento{
		Data:         request.Data,
		Horario:      request.Horario,
		Local:        request.Local,
		LocalMapsURL: request.LocalMapsURL,
		UpdatedAt:    time.Now(),
	}

	// Salva no arquivo
	fileContent, err := json.MarshalIndent(evento, "", "    ")
	if err != nil {
		log.Printf("Erro ao fazer marshal do JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao processar informações do evento: %v", err)})
		return
	}

	if err := os.WriteFile(eventoFilePath, fileContent, 0644); err != nil {
		log.Printf("Erro ao salvar arquivo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao salvar informações do evento: %v", err)})
		return
	}

	c.JSON(http.StatusOK, evento)
}
