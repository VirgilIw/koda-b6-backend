package service

import (
	"context"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUsers(ctx context.Context) ([]dto.Users, error) {

	datas, err := u.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	if len(datas) == 0 {
		return []dto.Users{}, nil
	}
	var result []dto.Users
	for _, v := range datas {
		result = append(result, dto.Users{
			Id:       v.Id,
			FullName: v.FullName,
			Email:    v.Email,
			Password: v.Password,
			Picture:  v.Picture,
			Phone:    v.Phone,
			Address:  v.Address,
			Role:     v.Role,
		})
	}
	return result, nil
}
