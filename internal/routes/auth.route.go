package routes

import (
	"github.com/MGT06/shortlink-app_backend/internal/handler"
	"github.com/MGT06/shortlink-app_backend/internal/repository"
	"github.com/MGT06/shortlink-app_backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func authRouter(router *gin.RouterGroup, db *pgxpool.Pool) {
	ar := repository.NewAuthRepo(db)
	as := service.NewAuthService(ar)
	ah := handler.NewAuthHandler(as)

	router.POST("register", ah.Register)
	router.POST("login", ah.Login)
}