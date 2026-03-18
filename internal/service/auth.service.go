package service

import (
	"context"
	"errors"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/lib"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (l *AuthService) AuthLogin(ctx context.Context, email, password string) string {

	user, err := l.repo.GetByEmail(ctx, email)

	if err != nil {
		return ""
	}

	if lib.VerifyPassword(password, user.Password) {
		token, err := lib.GenerateToken(user.Id)

		if err != nil {
			return ""
		}

		return token
	}

	return ""
}

func (l *AuthService) AuthRegister(ctx context.Context, req dto.AuthRegisterRequest) error {

	user, err := l.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	// kalau user ditemukan
	if user.Id != 0 {
		return errors.New("email already registered")
	}

	hashedPassword, err := lib.HashPassword(req.Password)

	if err != nil {
		return err
	}

	req.Password = hashedPassword

	_, err = l.repo.CreateUser(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
