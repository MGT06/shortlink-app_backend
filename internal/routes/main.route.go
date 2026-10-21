package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func MainRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	r := router.Group("api")
	
	authRouter(r, db)
}
