package dto

type CreateUserRequest struct {
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
