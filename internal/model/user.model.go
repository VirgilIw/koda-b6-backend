package model

import "time"

type UserModel struct {
	Id          int        `db:"id"`
	FullName    string     `db:"fullname"`
	Email       string     `db:"email"`
	Password    string     `db:"password"`
	Picture     *string    `db:"picture"`
	Phone       *string    `db:"phone"`
	Address     *string    `db:"address"`
	Role        *string    `db:"role"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
	LastLoginAt *time.Time `db:"lastlogin_at"`
}

type UserRegister struct {
	Id       int    `db:"id"`
	FullName string `db:"fullname"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
