package repository

import (
	"context"

	errormsg "github.com/MGT06/shortlink-app_backend/internal/error"
	"github.com/MGT06/shortlink-app_backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepo struct {
	db *pgxpool.Pool
}

func NewLinkRepo(db *pgxpool.Pool) *LinkRepo {
	return &LinkRepo{
		db: db,
	}
}

func (l *LinkRepo) CreateLink(ctx context.Context, body model.Links) error {
	query := "INSERT INTO links (user_id, original_url, slug) VALUES ($1, $2, $3)"
	args := []any{body.UserId, body.OriginaluUrl, body.Slug}

	cmt, err := l.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return errormsg.ErrNoRowAffected
	}

	return nil
}

func (l *LinkRepo) CheckLink(ctx context.Context, slug string) error {
	query := "SELECT EXISTS (SELECT 1 FROM links WHERE slug = $1)"
	args := []any{slug}

	var exists bool
	err := l.db.QueryRow(ctx, query, args...).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errormsg.ErrSlugTaken
	}

	return nil
}

func (l *LinkRepo) RedirectLink(ctx context.Context, slug string) (string, error) {
	query := "SELECT original_url FROM links WHERE slug = $1"
	args := []any{slug}

	var url string
	err := l.db.QueryRow(ctx, query, args...).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (l *LinkRepo) DeleteLink(ctx context.Context, linkId int) error {
	query := "UPDATE links SET delete_at = NOW() WHERE id = $1"
	args := []any{linkId}

	cmt, err := l.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return errormsg.ErrNoRowAffected
	}

	return nil
}

