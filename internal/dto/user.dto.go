package dto

import "time"

type CreateUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone,omitempty"`
	Address  string `json:"address,omitempty"`
	Picture  string `json:"picture,omitempty"`
	Role     string `json:"role,omitempty"`
}

type CreateUserResponse struct {
	Id        int        `json:"id"`
	FullName  string     `json:"full_name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone,omitempty"`
	Address   string     `json:"address,omitempty"`
	Picture   string     `json:"picture,omitempty"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	LastLogin *time.Time `json:"last_login,omitempty"`
}

type UpdateUserRequest struct {
	Id       int     `json:"id"`
	FullName string  `json:"fullname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Picture  *string `json:"picture"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
	Role     *string `json:"role"`
}

type Users struct {
	Id       int     `json:"id"`
	FullName string  `json:"fullname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Picture  *string `json:"picture"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
	Role     *string `json:"role"`
}
