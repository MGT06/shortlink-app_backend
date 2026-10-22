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

func authRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	r := router.Group("api")

	ar := repository.NewAuthRepo(db)
	as := service.NewAuthService(ar, rc)
	ah := handler.NewAuthHandler(as)

	r.POST("register", ah.Register)
	r.POST("login", ah.Login)
	r.POST("logout", middleware.CheckToken(rc), ah.Logout)
}