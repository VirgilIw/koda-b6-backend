package service

import (
	"context"
	"errors"
	"math/rand"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/lib"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type ForgotPwdService struct {
	fpwdRepo *repository.ForgotPwdRepository
	userRepo *repository.UserRepository
}

func NewForgotPwdService(fpwdrepo *repository.ForgotPwdRepository, userRepo *repository.UserRepository) *ForgotPwdService {
	return &ForgotPwdService{
		fpwdRepo: fpwdrepo,
		userRepo: userRepo,
	}
}

func (f *ForgotPwdService) RequestForgotPassword(ctx context.Context, req dto.ForgotPwdRequest) (dto.ForgotPwdResponse, error) {
	if req.Email == "" {
		return dto.ForgotPwdResponse{}, errors.New("email cannot be empty")
	}

	user, err := f.userRepo.GetByEmail(ctx, req.Email)

	if err != nil {
		return dto.ForgotPwdResponse{}, errors.New("email not found")
	}

	otp := rand.Intn(900000) + 100000

	forgotReq := dto.ForgotPwdRequest{
		Email: user.Email,
	}

	_, err = f.fpwdRepo.CreateForgotRequest(ctx, forgotReq, otp)
	if err != nil {
		return dto.ForgotPwdResponse{}, err
	}

	return dto.ForgotPwdResponse{
		Email:   user.Email,
		CodeOtp: otp,
	}, nil
}

func (f *ForgotPwdService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {

	_, err := f.fpwdRepo.GetDataByEmailAndCode(ctx, dto.ForgotPwdRequest{
		Email: req.Email,
	}, req.CodeOtp)

	if err != nil {
		return errors.New("invalid OTP")
	}

	user, err := f.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return errors.New("user not found")
	}

	hashedPassword, err := lib.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = hashedPassword

	err = f.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return errors.New("failed to update password")
	}

	err = f.fpwdRepo.DeleteDataByCode(ctx, dto.ForgotPwdRequest{
		Email: req.Email,
	}, req.CodeOtp)

	if err != nil {
		return errors.New("failed to delete OTP")
	}

	return nil
}
