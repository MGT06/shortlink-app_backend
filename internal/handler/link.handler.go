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

type LinkHandler struct {
	ls *service.LinkService
}

func NewLinkRepo(ls *service.LinkService) *LinkHandler {
	return &LinkHandler{
		ls: ls,
	}
}

// Create Link
//
// @Summary			Create a Short Link
// @Tags			link
// @Accept			json
// @Produce			json
// @Router			/api/links			[post]
// @Param			body	body		dto.CreateLinkRequest	true	"Link"
// @Security 		BearerToken
// @Success			201		{object}	dto.Response
// @Failure			400		{object}	dto.ErrorResponse
// @Failure			409		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (l LinkHandler) CreateLink(ctx *gin.Context) {
	id, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "a system error has occurred",
		})
		return
	}

	var body dto.CreateLinkRequest
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "invalid body",
		})
		return
	}

	if err := l.ls.CreateLink(ctx.Request.Context(), id.(int), body); err != nil {
		log.Println(err)
		if errors.Is(err, errormsg.ErrSlugTaken) {
			ctx.JSON(http.StatusConflict, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		}
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "a system error has occured",
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Link created successfully",
	})
}
