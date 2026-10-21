package dto

type Response struct {
	Success bool   `example:"true"`
	Message string `example:"success"`
	Results any    
}

type ErrorResponse struct {
	Success bool   `example:"false"`
	Message string `example:"error"`
	Results any    
}

