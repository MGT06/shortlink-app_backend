package repository

import (
	"context"
	"fmt"

	errormsg "github.com/MGT06/shortlink-app_backend/internal/error"
	"github.com/MGT06/shortlink-app_backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type LinkRepo struct{}

func NewLinkRepo() *LinkRepo {
	return &LinkRepo{}
}

func (l *LinkRepo) CreateLink(ctx context.Context, db DBTX, body model.Links) error {
	query := "INSERT INTO links (user_id, original_url, slug) VALUES ($1, $2, $3)"
	args := []any{body.UserId, body.OriginaluUrl, body.Slug}

	cmt, err := db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return errormsg.ErrNoRowAffected
	}

	return nil
}

func (l *LinkRepo) CheckLink(ctx context.Context, db DBTX, slug string) error {
	query := "SELECT EXISTS (SELECT 1 FROM links WHERE slug = $1)"
	args := []any{slug}

	var exists bool
	err := db.QueryRow(ctx, query, args...).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errormsg.ErrSlugTaken
	}

	return nil
}

func (l *LinkRepo) RedirectLink(ctx context.Context, db DBTX, slug string) (string, error) {
	query := "SELECT original_url FROM links WHERE slug = $1"
	args := []any{slug}

	var url string
	err := db.QueryRow(ctx, query, args...).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (l *LinkRepo) GetAllLink(ctx context.Context, db DBTX, userId int, search, sort string, limit, offset int) ([]model.Links, error) {
	query := fmt.Sprintf("SELECT id, original_url, slug, created_at FROM links WHERE user_id = $1 AND delete_at IS NULL AND (original_url ILIKE $2 OR slug ILIKE $2) ORDER BY created_at %s LIMIT $3 OFFSET $4", sort)

	args := []any{userId, "%" + search + "%", limit, offset}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("dfadf", err)
	}
	defer rows.Close()

	links := make([]model.Links, 0, limit)
	for rows.Next() {
		var link model.Links
		if err := rows.Scan(&link.LinkId, &link.OriginaluUrl, &link.Slug, &link.Created_at); err != nil {
			return nil, fmt.Errorf("dfadf", err)
		}
		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("dfadf", err)
	}

	return links, nil
}

func (l *LinkRepo) GetAllLinkMeta(ctx context.Context, db DBTX, userId int) (int, error) {
	query := "SELECT COUNT(*) FROM links WHERE user_id = $1 AND delete_at IS NULL"
	args := []any{userId}

	var total int
	if err := db.QueryRow(ctx, query, args...).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (l *LinkRepo) DeleteLink(ctx context.Context, db DBTX, linkId int) error {
	query := "UPDATE links SET delete_at = NOW() WHERE id = $1"
	args := []any{linkId}

	cmt, err := db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmt.RowsAffected() == 0 {
		return errormsg.ErrNoRowAffected
	}

	return nil
}
