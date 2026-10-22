package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

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

// Register
//
// @Summary			Register
// @Description		Register Account
// @Tags			auth
// @Accept			json
// @Produce			json
// @Router			/api/register	[post]
// @Param			newAccount	body	dto.RegisterRequest		true	"Register"
// @Success			200		{object}	dto.Response
// @Failure			400		{object}	dto.ErrorResponse
// @Failure			409		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (a *AuthHandler) Register(ctx *gin.Context) {
	var newAccount dto.RegisterRequest
	if err := ctx.ShouldBindWith(&newAccount, binding.JSON); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "invalid body",
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

// Login
//
// @Summary			Login
// @Description		Login using email and password
// @Tags			auth
// @Accept			json
// @Produce			json
// @Router			/api/login	[post]
// @Param			account	body	dto.LoginRequest	true	"Login"
// @Success			200		{object}	dto.Response
// @Failure			400		{object}	dto.ErrorResponse
// @Failure			401		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (a *AuthHandler) Login(ctx *gin.Context) {
	var account dto.LoginRequest
	if err := ctx.ShouldBindWith(&account, binding.JSON); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Invalid email format",
		})
		return
	}

	token, err := a.as.Login(ctx.Request.Context(), account)
	if err != nil {
		log.Println(err)
		if errors.Is(err, errormsg.ErrEmailPasswordIncorrect) {
			ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Login Success",
		Results: dto.LoginResponse{
			Token: token,
		},
	})
}

// Logout
//
// @Summary			Logut
// @Description		User Logout
// @Tags			auth
// @Accept			json
// @Produce			json
// @Router			/api/logout 	[post]
// @Security 		BearerToken
// @Success			200		{object}	dto.Response
// @Failure			400		{object}	dto.ErrorResponse
// @Failure			401		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (a *AuthHandler) Logout(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	jwtID, exist := ctx.Get("jti")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	exp, exist := ctx.Get("exp")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	expTime, ok := exp.(time.Time)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ttl := time.Until(expTime)

	if err := a.as.Logout(ctx, userId.(int), jwtID.(string), ttl); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "A system error has occurred",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, dto.Response{
		Success: true,
		Message: "logout success",
	})
}