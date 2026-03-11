package model

import "time"

type ForgotPassword struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	CodeOtp   int       `db:"code_otp"`
	CreatedAt time.Time `db:"created_at"`
}
