package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	_ "github.com/MGT06/shortlink-app_backend/docs"
	"github.com/MGT06/shortlink-app_backend/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MainRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	router.Use(middleware.Cors)

	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRouter(router, db, rc)
	linkRouter(router, db, rc)
}
