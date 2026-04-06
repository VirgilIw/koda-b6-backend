package service

import (
	"context"
	"errors"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/lib"
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

func (u *UserService) GetUserById(ctx context.Context, id int) (dto.Users, error) {
	data, err := u.repo.GetById(ctx, id)

	if err != nil {
		return dto.Users{}, err
	}

	var result dto.Users

	result = dto.Users{
		Id:       data.Id,
		FullName: data.FullName,
		Email:    data.Email,
		Password: data.Password,
		Picture:  data.Picture,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
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
	if req.Role != "" {
		user.Role = req.Role
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return model.UserModel{}, errors.New("failed to update user")
	}

	return user, nil
}

func (u *UserService) DeleteUser(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid user id")
	}

	return u.repo.DeleteUser(ctx, id)
}

func (u *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	// Mapping DTO request → Model
	reqUser := dto.AuthRegisterRequest{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}

	// Hash password
	hashed, err := lib.HashPassword(reqUser.Password)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}
	reqUser.Password = hashed

	newUser, err := u.repo.CreateUser(ctx, reqUser)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	// Mapping Model → DTO response
	resp := dto.CreateUserResponse{
		Id:       newUser.Id,
		FullName: newUser.FullName,
		Email:    newUser.Email,
	}

	return resp, nil
}
