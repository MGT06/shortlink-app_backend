package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Access-Control-Max-Age", "3600")

	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}