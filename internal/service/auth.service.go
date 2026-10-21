package service

import (
	"context"

	"github.com/MGT06/shortlink-app_backend/internal/dto"
	errormsg "github.com/MGT06/shortlink-app_backend/internal/error"
	"github.com/MGT06/shortlink-app_backend/internal/model"
	"github.com/MGT06/shortlink-app_backend/internal/repository"
	"github.com/MGT06/shortlink-app_backend/pkg"
)

type AuthService struct {
	ar *repository.AuthRepo
}

func NewAuthService(ar *repository.AuthRepo) *AuthService {
	return &AuthService{
		ar: ar,
	}
}

func (a *AuthService) Register(ctx context.Context, body dto.RegisterRequest) error {
	_, err := a.ar.FindAccount(ctx, body.Email)
	if err != nil {
		return err
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
		return "", errormsg.ErrEmailPasswordIncorrect
	}

	if err := pkg.Compare(body.Password, acc.Password); err != nil {
		return "", errormsg.ErrEmailPasswordIncorrect
	}

	claims := pkg.NewJWTClaims(acc.Id)
	return claims.GenToken()
}
