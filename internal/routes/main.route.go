package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	_ "github.com/MGT06/shortlink-app_backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MainRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	r := router.Group("api")
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRouter(r, db)
	linkRouter(r, db, rc)
}
