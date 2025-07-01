package models

import (
	"time"
)

// Confirmacao representa uma confirmação de presença
type Confirmacao struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	Nome               string    `json:"nome" gorm:"not null"`
	QuantidadeAdultos  int       `json:"quantidade_adultos" gorm:"not null"`
	QuantidadeCriancas int       `json:"quantidade_criancas" gorm:"default:0"`
	CriadaEm           time.Time `json:"criada_em" gorm:"autoCreateTime"`
}

// ConfirmacaoPublica representa a visão pública de uma confirmação
type ConfirmacaoPublica struct {
	Nome               string `json:"nome"`
	QuantidadeAdultos  int    `json:"quantidade_adultos"`
	QuantidadeCriancas int    `json:"quantidade_criancas"`
}

// ToPublic converte uma Confirmacao para ConfirmacaoPublica
func (c *Confirmacao) ToPublic() ConfirmacaoPublica {
	return ConfirmacaoPublica{
		Nome:               c.Nome,
		QuantidadeAdultos:  c.QuantidadeAdultos,
		QuantidadeCriancas: c.QuantidadeCriancas,
	}
}

// CreateConfirmacaoRequest representa o request para criar uma confirmação
type CreateConfirmacaoRequest struct {
	Nome               string `json:"nome" binding:"required"`
	QuantidadeAdultos  int    `json:"quantidade_adultos" binding:"required,min=1"`
	QuantidadeCriancas int    `json:"quantidade_criancas" binding:"min=0"`
}
