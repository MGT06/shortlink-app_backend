package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/MGT06/shortlink-app_backend/internal/dto"
	errormsg "github.com/MGT06/shortlink-app_backend/internal/error"
	"github.com/MGT06/shortlink-app_backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5"
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

func (l *LinkHandler) RedirectLink(ctx *gin.Context) {
	slug := ctx.Param("slug")

	url, err := l.ls.RedirectLink(ctx.Request.Context(), slug)
	if err != nil {
		log.Println(err)
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
				Success: false,
				Message: errormsg.ErrLinkNotFound.Error(),
			})
		}
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "a system error has occured",
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url)
}

// Get All Links
//
// @Summary			Get All Short Link
// @Tags			link
// @Accept			json
// @Produce			json
// @Router			/api/links			[get]
// @Param			query	query		dto.GetLinksQuery	false	"query param"
// @Security 		BearerToken
// @Success			201		{object}	dto.Response
// @Failure			400		{object}	dto.ErrorResponse
// @Failure			409		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (l *LinkHandler) GetAllLinks(ctx *gin.Context) {
	id, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "a system error has occurred",
		})
		return
	}

	var query dto.GetLinksQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "a system error has occured",
		})
		return
	}

	res, err := l.ls.GetAllLinks(ctx, id.(int), query)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "a system error has occured",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success Get Links",
		Results: res,
	})
}

// Delete Link
//
// @Summary			Delete a Short Link
// @Tags			link
// @Accept			json
// @Produce			json
// @Router			/api/links/{id}			[delete]
// @Param			linkId	path		string	true	"id link"
// @Security 		BearerToken
// @Success			200		{object}	dto.Response
// @Failure			404		{object}	dto.ErrorResponse
// @Failure			422		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func (l *LinkHandler) DeleteLink(ctx *gin.Context) {
	linkId := ctx.Param("id")

	id, err := strconv.Atoi(linkId)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Success: false,
			Message: "invalid param",
		})
		return
	}

	if err := l.ls.DeleteLink(ctx.Request.Context(), id); err != nil {
		log.Println(err)
		if errors.Is(err, errormsg.ErrNoRowAffected) {
			ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
				Success: false,
				Message: "link not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "a system error has occured",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Link Deleted",
	})

}
