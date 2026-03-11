package dto

type ForgotPwdRequest struct {
	Email   string `json:"email" validate:"required,email"`
	CodeOtp int    `json:"code_otp" validate:"required"`
}
