package model

import "time"

type Transaction struct {
	ID              int        `db:"id"`
	TransactionCode string     `db:"transaction_code"`
	DeliveryMethod  string     `db:"delivery_method"`
	FullName        string     `db:"full_name"`
	Email           string     `db:"email"`
	Address         string     `db:"address"`
	SubtotalPrice   int        `db:"subtotal_price"`
	TotalPrice      int        `db:"total_price"`
	DeliveryFee     int        `db:"delivery_fee"`
	Tax             int        `db:"tax"`
	CouponID        *int       `db:"coupon_id"`
	Status          string     `db:"status"`
	PaymentMethod   string     `db:"payment_method"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at"`
}
