package dto

type CreateLinkRequest struct {
	OriginaluUrl string `json:"original_url" binding:"required"`
	Slug         string `json:"slug" binding:"omitempty,len=6"`
}
