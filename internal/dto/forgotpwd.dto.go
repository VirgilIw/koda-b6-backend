package dto

type ForgotPwdRequest struct {
	Email string `json:"email" validate:"required,email"`
	// CodeOtp int    `json:"code_otp,omitempty" validate:"required"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" example:"user@email.com"`
	CodeOtp     int    `json:"code_otp" example:"123456"`
	NewPassword string `json:"new_password" example:"newpassword123"`
}
