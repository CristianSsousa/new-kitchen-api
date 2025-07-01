package models

// PublicModel define a interface para modelos que têm uma visão pública
type PublicModel interface {
	ToPublic() interface{}
}

// AdminResponse define a estrutura base para respostas administrativas
type AdminResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse define a estrutura para respostas de erro
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// PaginatedResponse define a estrutura para respostas paginadas
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	TotalPages int         `json:"total_pages"`
}
