package models

import (
	"time"
)

// Mensagem representa uma mensagem dos convidados
type Mensagem struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Nome       string    `json:"nome" gorm:"not null"`
	Mensagem   string    `json:"mensagem" gorm:"not null"`
	Aprovada   bool      `json:"aprovada" gorm:"default:false"`
	CriadaEm   time.Time `json:"criada_em" gorm:"autoCreateTime"`
	AprovadaEm time.Time `json:"aprovada_em,omitempty"`
}

// MensagemPublica representa a visão pública de uma mensagem
type MensagemPublica struct {
	Nome     string `json:"nome"`
	Mensagem string `json:"mensagem"`
}

// ToPublic converte uma Mensagem para MensagemPublica
func (m *Mensagem) ToPublic() MensagemPublica {
	return MensagemPublica{
		Nome:     m.Nome,
		Mensagem: m.Mensagem,
	}
}

// CreateMensagemRequest representa o request para criar uma mensagem
type CreateMensagemRequest struct {
	Nome     string `json:"nome" binding:"required"`
	Mensagem string `json:"mensagem" binding:"required"`
}
