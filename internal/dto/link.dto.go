package dto

import "time"

type CreateLinkRequest struct {
	OriginaluUrl string `json:"original_url" binding:"required"`
	Slug         string `json:"slug" binding:"omitempty,len=6"`
}

type GetLinksQuery struct {
	Search string `form:"search"`
	Sort   string `form:"sort" binding:"omitempty,oneof=asc desc"`
	Page   int    `form:"page" binding:"omitempty,min=1"`
	Limit  int    `form:"limit" binding:"omitempty,min=5,max=10"`
}

type LinkItem struct {
	Id          int       `json:"id"`
	OriginalURL string    `json:"original_url"`
	Slug        string    `json:"slug"`
	Click       int       `json:"click"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetLinksData struct {
	Links      []LinkItem     `json:"links"`
	Pagination PaginationMeta `json:"pagination"`
}
