package service

import (
	"context"
	"errors"
	"log"
	"math/rand"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type ForgotPwdService struct {
	fpwdRepo *repository.ForgotPwdRepository
	userRepo *repository.UserRepository
}

func NewForgotService(fpwdrepo *repository.ForgotPwdRepository, userRepo *repository.UserRepository) *ForgotPwdService {
	return &ForgotPwdService{
		fpwdRepo: fpwdrepo,
		userRepo: userRepo,
	}
}

func (f *ForgotPwdService) RequestForgotPassword(ctx context.Context, req dto.ForgotPwdRequest) (dto.ForgotPwdResponse, error) {
	// pastikan email ada di repo user
	user, err := f.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return dto.ForgotPwdResponse{}, errors.New("email not found")
	}

	// generate code
	otp := rand.Intn(900000) + 100000
	log.Println("OTP:", otp)

	forgotReq := dto.ForgotPwdRequest{
		Email:   user.Email,
		CodeOtp: otp,
	}

	// masukan (insert) email dan otp dari user ke repo
	_, err = f.fpwdRepo.CreateForgotRequest(ctx, forgotReq)
	if err != nil {
		return dto.ForgotPwdResponse{}, err
	}

	var forgotResp dto.ForgotPwdResponse

	forgotResp = dto.ForgotPwdResponse{
		Email:   user.Email,
		CodeOtp: otp,
	}

	return forgotResp, nil
}

func (f *ForgotPwdService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {

	// verifikasi OTP valid atau gak
	forgotData, err := f.fpwdRepo.GetDataByEmailAndCode(ctx, dto.ForgotPwdRequest{
		Email:   req.Email,
		CodeOtp: req.CodeOtp,
	})

	if err != nil {
		return errors.New("invalid OTP")
	}

	// ambil data user
	user, err := f.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return errors.New("user not found")
	}

	// update password
	user.Password = req.NewPassword
	err = f.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return errors.New("failed to update password")
	}

	// hapus OTP
	err = f.fpwdRepo.DeleteDataByCode(ctx, dto.ForgotPwdRequest{
		Email:   req.Email,
		CodeOtp: forgotData.CodeOtp,
	})

	if err != nil {
		return errors.New("failed to delete OTP")
	}

	return nil
}
