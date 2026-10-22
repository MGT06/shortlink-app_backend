package routes

import (
	"github.com/MGT06/shortlink-app_backend/internal/handler"
	"github.com/MGT06/shortlink-app_backend/internal/repository"
	"github.com/MGT06/shortlink-app_backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func authRouter(router *gin.Engine, db *pgxpool.Pool) {
	r := router.Group("api")

	ar := repository.NewAuthRepo(db)
	as := service.NewAuthService(ar)
	ah := handler.NewAuthHandler(as)

	r.POST("register", ah.Register)
	r.POST("login", ah.Login)
}