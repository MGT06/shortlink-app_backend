package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MGT06/shortlink-app_backend/internal/dto"
	errormsg "github.com/MGT06/shortlink-app_backend/internal/error"
	"github.com/MGT06/shortlink-app_backend/internal/model"
	"github.com/MGT06/shortlink-app_backend/internal/repository"
	"github.com/MGT06/shortlink-app_backend/internal/utils"
	"github.com/MGT06/shortlink-app_backend/pkg"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	ar *repository.AuthRepo
	rc *redis.Client
}

func NewAuthService(ar *repository.AuthRepo, rc *redis.Client) *AuthService {
	return &AuthService{
		ar: ar,
		rc: rc,
	}
}

func (a *AuthService) Register(ctx context.Context, body dto.RegisterRequest) error {
	_, err := a.ar.FindAccount(ctx, body.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return errormsg.ErrEmailExist
	}

	hash := pkg.NewHashConfig().GenHash(body.Password)

	if err := a.ar.Register(ctx, model.Users{
		Name:     body.Name,
		Email:    body.Email,
		Password: hash,
	}); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) Login(ctx context.Context, body dto.LoginRequest) (string, error) {
	acc, err := a.ar.FindAccount(ctx, body.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errormsg.ErrEmailPasswordIncorrect
		}
		return "", err
	}

	if err := pkg.Compare(body.Password, acc.Password); err != nil {
		return "", errormsg.ErrEmailPasswordIncorrect
	}

	claims := pkg.NewJWTClaims(acc.Id)
	return claims.GenToken()
}

func (a *AuthService) Logout(ctx context.Context, userId int, JTI string, Expired time.Duration) error {
	key := fmt.Sprintf("shortlink:tokenBlacklist:%s", JTI)

	if err := utils.SetToRedis(ctx, a.rc, key, userId, Expired); err != nil {
		return err
	}

	return nil
}