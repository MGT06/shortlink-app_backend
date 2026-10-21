package repository

import (
	"context"
	"errors"

	errormsg "github.com/MGT06/shortlink-app_backend/internal/error"
	"github.com/MGT06/shortlink-app_backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (a *AuthRepo) Register(ctx context.Context, body model.Users) error {
	query := "INSERT INTO accounts (full_name, email, password_hash) VALUES ($1, $2, $3)"
	args := []any{body.Name, body.Email, body.Password}

	cmt, err := a.db.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
			return errormsg.ErrEmailExist
		}
		return err
	}

	if cmt.RowsAffected() == 0 {
		return errormsg.ErrNoRowAffected
	}

	return nil
}

func (a *AuthRepo) FindAccount(ctx context.Context, email string) (model.Users, error) {
	query := "SELECT id, full_name, password from accounts WHERE email=$1"
	args := []any{email}

	var data model.Users
	if err := a.db.QueryRow(ctx, query, args...).Scan(&data.Id, &data.Name, &data.Password); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return model.Users{}, nil
		}
		return model.Users{}, err
	}

	return data, nil
}