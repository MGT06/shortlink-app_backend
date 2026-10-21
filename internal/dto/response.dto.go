package dto

type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success"`
	Results any    `json:"results"`
}

type PaginationMeta struct {
	Page        int `json:"page"`
	ActiveLinks int `json:"active_links"`
	TotalPages  int `json:"total_pages"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"error"`
}
