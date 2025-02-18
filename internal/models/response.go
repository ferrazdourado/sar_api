package models

// APIResponse representa a estrutura padrão de resposta da API
type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
    Message string      `json:"message,omitempty"`
}

// PaginatedResponse representa uma resposta paginada
type PaginatedResponse struct {
    Items      interface{} `json:"items"`
    Total      int64      `json:"total"`
    Page       int        `json:"page"`
    PageSize   int        `json:"page_size"`
    TotalPages int        `json:"total_pages"`
}