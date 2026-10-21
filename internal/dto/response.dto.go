package dto

type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success"`
	Results any    `json:"results"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"error"`
	Results any    `json:"results"`
}

