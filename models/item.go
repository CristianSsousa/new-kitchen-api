package models

import (
	"time"
)

// Item representa um item da lista de presentes
type Item struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Nome         string    `json:"nome" gorm:"not null"`
	Descricao    string    `json:"descricao"`
	Categoria    string    `json:"categoria"`
	Preco        float64   `json:"preco"`
	ImagemURL    string    `json:"imagem_url"`
	LinkURL      string    `json:"link_url,omitempty"`
	Resgatado    bool      `json:"resgatado" gorm:"default:false"`
	ResgatadoPor string    `json:"resgatado_por,omitempty"`
	ResgatadoEm  time.Time `json:"resgatado_em,omitempty"`
	CriadoEm     time.Time `json:"criado_em" gorm:"autoCreateTime"`
	AtualizadoEm time.Time `json:"atualizado_em" gorm:"autoUpdateTime"`
}

// ItemPublico representa a visão pública de um item
type ItemPublico struct {
	ID        uint    `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Categoria string  `json:"categoria"`
	Preco     float64 `json:"preco"`
	ImagemURL string  `json:"imagem_url"`
	LinkURL   string  `json:"link_url,omitempty"`
	Resgatado bool    `json:"resgatado"`
}

// ToPublic converte um Item para ItemPublico
func (i *Item) ToPublic() ItemPublico {
	return ItemPublico{
		ID:        i.ID,
		Nome:      i.Nome,
		Descricao: i.Descricao,
		Categoria: i.Categoria,
		Preco:     i.Preco,
		ImagemURL: i.ImagemURL,
		LinkURL:   i.LinkURL,
		Resgatado: i.Resgatado,
	}
}

// CreateItemRequest representa o request para criar um item
type CreateItemRequest struct {
	Nome      string  `json:"nome" binding:"required"`
	Descricao string  `json:"descricao"`
	Categoria string  `json:"categoria"`
	Preco     float64 `json:"preco"`
	ImagemURL string  `json:"imagem_url"`
	LinkURL   string  `json:"link_url,omitempty"`
}

// ResgatarItemRequest representa o request para resgatar um item
type ResgatarItemRequest struct {
	Nome string `json:"nome" binding:"required"`
}
