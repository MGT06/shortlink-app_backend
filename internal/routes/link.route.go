package routes

import (
	"github.com/MGT06/shortlink-app_backend/internal/handler"
	"github.com/MGT06/shortlink-app_backend/internal/middleware"
	"github.com/MGT06/shortlink-app_backend/internal/repository"
	"github.com/MGT06/shortlink-app_backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func linkRouter(router *gin.RouterGroup, db *pgxpool.Pool, rc *redis.Client) {
	lr := repository.NewLinkRepo(db)
	ls := service.NewLinkRepo(lr)
	lh := handler.NewLinkRepo(ls)

	router.POST("links",middleware.CheckToken(rc), lh.CreateLink)
}