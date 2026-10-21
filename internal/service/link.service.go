package service

import (
	"context"
	"math"
	"strings"

	"github.com/MGT06/shortlink-app_backend/internal/dto"
	"github.com/MGT06/shortlink-app_backend/internal/model"
	"github.com/MGT06/shortlink-app_backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.step.sm/crypto/randutil"
)

type LinkService struct {
	lr *repository.LinkRepo
	db *pgxpool.Pool
}

func NewLinkRepo(lr *repository.LinkRepo, db *pgxpool.Pool) *LinkService {
	return &LinkService{
		lr: lr,
		db: db,
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

	if err := l.lr.CheckLink(ctx, l.db, slug); err != nil {
		return err
	}

	if err := l.lr.CreateLink(ctx, l.db, model.Links{
		UserId:       userId,
		OriginaluUrl: body.OriginaluUrl,
		Slug:         slug,
	}); err != nil {
		return err
	}

	return nil
}

func (l *LinkService) RedirectLink(ctx context.Context, slug string) (string, error) {
	url, err := l.lr.RedirectLink(ctx, l.db, slug)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (l *LinkService) GetAllLinks(ctx context.Context, userId int, q dto.GetLinksQuery) (dto.GetLinksData, error) {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.Limit < 1 {
		q.Limit = 10
	}
	if q.Sort == "" {
		q.Sort = "desc"
	}

	q.Search = strings.TrimSpace(q.Search)

	offset := (q.Page - 1) * q.Limit

	links, err := l.lr.GetAllLink(ctx, l.db, userId, q.Search, q.Sort, q.Limit, offset)
	if err != nil {
		return dto.GetLinksData{}, err
	}

	total, err := l.lr.GetAllLinkMeta(ctx, l.db, userId)
	if err != nil {
		return dto.GetLinksData{}, err
	}

	data := make([]dto.LinkItem, 0, len(links))
	for _, link := range links {
		data = append(data, dto.LinkItem{
			Id:          link.LinkId,
			OriginalURL: link.OriginaluUrl,
			Slug:        link.Slug,
			CreatedAt:   link.Created_at,
		})
	}

	return dto.GetLinksData{
		Links: data,
		Pagination: dto.PaginationMeta{
			Page:        q.Page,
			ActiveLinks: total,
			TotalPages:  int(math.Ceil(float64(total) / float64(q.Limit))),
		},
	}, nil
}

func (l *LinkService) DeleteLink(ctx context.Context, linkId int) error {
	if err := l.lr.DeleteLink(ctx, l.db, linkId); err != nil {
		return err
	}

	return nil
}
