package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/MGT06/shortlink-app_backend/internal/dto"
	errormsg "github.com/MGT06/shortlink-app_backend/internal/error"
	"github.com/MGT06/shortlink-app_backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthHandler struct {
	as *service.AuthService
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{
		as: as,
	}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var newAccount dto.RegisterRequest
	if err := ctx.ShouldBindWith(&newAccount, binding.JSON); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	if err := a.as.Register(ctx.Request.Context(), newAccount); err != nil {
		log.Println(err)
		if errors.Is(err, errormsg.ErrEmailExist) {
			ctx.JSON(http.StatusConflict, dto.ErrorResponse{
				Success: false,
				Message: errormsg.ErrEmailExist.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Account created successfully",
	})
}