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

func linkRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	r := router.Group("api")

	lr := repository.NewLinkRepo()
	ls := service.NewLinkRepo(lr, db)
	lh := handler.NewLinkRepo(ls)

	r.POST("links",middleware.CheckToken(rc), lh.CreateLink)
	r.GET("links",middleware.CheckToken(rc), lh.GetAllLinks)
	r.DELETE("links/:id", middleware.CheckToken(rc), lh.DeleteLink)
	router.GET(":slug", lh.RedirectLink)
}