package service

import (
	"context"
	"errors"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
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

// UpdateProfile updates user profile based on UpdateUserRequest
func (u *UserService) UpdateProfile(ctx context.Context, req dto.UpdateUserRequest) (model.UserModel, error) {
	user, err := u.repo.GetById(ctx, req.Id)
	if err != nil {
		return model.UserModel{}, errors.New("user not found")
	}

	// Update fields yang diisi
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Picture != nil {
		user.Picture = req.Picture
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.Address != nil {
		user.Address = req.Address
	}
	if req.Role != nil {
		user.Role = req.Role
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return model.UserModel{}, errors.New("failed to update user")
	}

	return user, nil
}
