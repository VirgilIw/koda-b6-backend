package dto

type AuthLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthRegisterRequest struct {
	FullName string `json:"fullname" binding:"required,min=3,max=100" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@email.com"`
	Password string `json:"password" binding:"required,min=6,max=100" example:"password123"`
}
