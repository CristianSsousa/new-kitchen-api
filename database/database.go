package database

import (
	"cha-casa-nova-backend/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	fmt.Println("Database connected successfully")

	// Auto Migrate
	err = DB.AutoMigrate(
		&models.Item{},
		&models.Mensagem{},
		&models.Confirmacao{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	fmt.Println("Database migrated successfully")

	// Criar dados de exemplo se o banco estiver vazio
	var count int64
	DB.Model(&models.Item{}).Count(&count)
	if count == 0 {
		seedData()
	}
}

func seedData() {
	// Itens para cozinha
	itensParaCozinha := []models.Item{
		{
			Nome:      "Jogo de Panelas",
			Descricao: "Conjunto de panelas antiaderentes com 5 peças",
			Categoria: "Cozinha",
			Preco:     299.90,
			ImagemURL: "https://example.com/panelas.jpg",
		},
		{
			Nome:      "Liquidificador",
			Descricao: "Liquidificador 1000W com 5 velocidades",
			Categoria: "Cozinha",
			Preco:     199.90,
			ImagemURL: "https://example.com/liquidificador.jpg",
		},
		{
			Nome:      "Jogo de Talheres",
			Descricao: "Kit com 24 peças em aço inox",
			Categoria: "Cozinha",
			Preco:     149.90,
			ImagemURL: "https://example.com/talheres.jpg",
		},
	}

	// Itens para sala
	itensParaSala := []models.Item{
		{
			Nome:      "Almofadas Decorativas",
			Descricao: "Kit com 4 almofadas decorativas",
			Categoria: "Sala",
			Preco:     129.90,
			ImagemURL: "https://example.com/almofadas.jpg",
		},
		{
			Nome:      "Tapete",
			Descricao: "Tapete 2x1.5m para sala",
			Categoria: "Sala",
			Preco:     199.90,
			ImagemURL: "https://example.com/tapete.jpg",
		},
	}

	// Itens para quarto
	itensParaQuarto := []models.Item{
		{
			Nome:      "Jogo de Cama",
			Descricao: "Kit com lençol, fronhas e edredom queen size",
			Categoria: "Quarto",
			Preco:     249.90,
			ImagemURL: "https://example.com/jogocama.jpg",
		},
		{
			Nome:      "Travesseiros",
			Descricao: "Par de travesseiros antialérgicos",
			Categoria: "Quarto",
			Preco:     99.90,
			ImagemURL: "https://example.com/travesseiros.jpg",
		},
	}

	// Itens para banheiro
	itensParaBanheiro := []models.Item{
		{
			Nome:      "Jogo de Toalhas",
			Descricao: "Kit com 4 toalhas de banho e 4 de rosto",
			Categoria: "Banheiro",
			Preco:     159.90,
			ImagemURL: "https://example.com/toalhas.jpg",
		},
		{
			Nome:      "Tapete de Banheiro",
			Descricao: "Jogo com 3 tapetes antiderrapantes",
			Categoria: "Banheiro",
			Preco:     79.90,
			ImagemURL: "https://example.com/tapetebanheiro.jpg",
		},
	}

	// Inserir todos os itens
	for _, item := range append(append(append(itensParaCozinha, itensParaSala...), itensParaQuarto...), itensParaBanheiro...) {
		if err := DB.Create(&item).Error; err != nil {
			log.Printf("Erro ao criar item %s: %v", item.Nome, err)
		}
	}

	log.Println("✨ Dados de exemplo criados com sucesso!")
}
