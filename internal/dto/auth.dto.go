package dto

type AuthLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthRegisterRequest struct {
	FullName string `json:"fullname" example:"John Doe"`
	Email    string `json:"email" example:"john@email.com"`
	Password string `json:"password" example:"password123"`
}
