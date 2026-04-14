package service

import (
	"context"
	"errors"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/lib"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
			Id:        v.Id,
			FullName:  v.FullName,
			Email:     v.Email,
			Password:  v.Password,
			Picture:   v.Picture,
			Phone:     v.Phone,
			Address:   v.Address,
			Role:      v.Role,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
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
		Id:        data.Id,
		FullName:  data.FullName,
		Email:     data.Email,
		Password:  data.Password,
		Picture:   data.Picture,
		Phone:     data.Phone,
		Address:   data.Address,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return result, nil
}

// UpdateProfile updates user profile based on UpdateUserRequest
func (u *UserService) UpdateProfile(ctx context.Context, userId int, req dto.UpdateUserRequest, picture *string) (model.UserModel, error) {

	user, err := u.repo.GetById(ctx, userId)
	if err != nil {
		return model.UserModel{}, errors.New("user not found")
	}

	if req.FullName != nil {
		user.FullName = *req.FullName
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	if req.Password != nil {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		user.Password = string(hashed)
	}

	if picture != nil {
		user.Picture = picture
	}

	if req.Phone != nil {
		user.Phone = req.Phone
	}

	if req.Address != nil {
		user.Address = req.Address
	}

	err = u.repo.UpdateUser(ctx, user)
	if err != nil {
		return model.UserModel{}, err
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
