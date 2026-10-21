package service

import (
	"context"

	"github.com/MGT06/shortlink-app_backend/internal/dto"
	"github.com/MGT06/shortlink-app_backend/internal/model"
	"github.com/MGT06/shortlink-app_backend/internal/repository"
	"go.step.sm/crypto/randutil"
)

type LinkService struct {
	lr *repository.LinkRepo
}

func NewLinkRepo(lr *repository.LinkRepo) *LinkService {
	return &LinkService{
		lr: lr,
	}
}

func (l LinkService) CreateLink(ctx context.Context, userId int, body dto.CreateLinkRequest) error {
	randomSlug, err := randutil.Alphanumeric(6)
	if err != nil {
		return err
	}

	var slug string
	if body.Slug != "" {
		slug = body.Slug
	} else if body.Slug == "" {
		slug = randomSlug
	}

	if err := l.lr.CheckLink(ctx, slug); err != nil {
		return err
	}

	if err := l.lr.CreateLink(ctx, model.Links{
		UserId:       userId,
		OriginaluUrl: body.OriginaluUrl,
		Slug:         slug,
	}); err != nil {
		return err
	}

	return nil
}

func (l *LinkService) RedirectLink(ctx context.Context, slug string) (string, error) {
	url, err := l.lr.RedirectLink(ctx, slug)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (l *LinkService) DeleteLink(ctx context.Context, linkId int) error {
	if err := l.lr.DeleteLink(ctx, linkId); err != nil {
		return err
	}

	return nil
}
