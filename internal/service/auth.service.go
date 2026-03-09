package service

import (
	"context"
	"fmt"

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
	fmt.Println("email:", email)
	fmt.Println("password:", password)
	fmt.Println("hash:", user.Password)

	ok := lib.VerifyPassword(password, user.Password)
	fmt.Println("verify:", ok)
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
