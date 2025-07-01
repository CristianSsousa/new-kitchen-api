package models

import (
	"time"
)

// Evento representa as informações do evento
type Evento struct {
	Data         string    `json:"data"`
	Horario      string    `json:"horario"`
	Local        string    `json:"local"`
	LocalMapsURL string    `json:"local_maps_url,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// EventoPublico representa a visão pública do evento
type EventoPublico struct {
	Data         string `json:"data"`
	Horario      string `json:"horario"`
	Local        string `json:"local"`
	LocalMapsURL string `json:"local_maps_url,omitempty"`
}

// ToPublic converte um Evento para EventoPublico
func (e *Evento) ToPublic() EventoPublico {
	return EventoPublico{
		Data:         e.Data,
		Horario:      e.Horario,
		Local:        e.Local,
		LocalMapsURL: e.LocalMapsURL,
	}
}

// UpdateEventoRequest representa o request para atualizar informações do evento
type UpdateEventoRequest struct {
	Data         string `json:"data" binding:"required"`
	Horario      string `json:"horario" binding:"required"`
	Local        string `json:"local" binding:"required"`
	LocalMapsURL string `json:"local_maps_url,omitempty"`
}
