package model

import "time"

type Transaction struct {
	ID              int       `db:"id"`
	UserID          int       `db:"user_id"`
	TransactionCode string    `db:"transaction_code"`
	DeliveryMethod  string    `db:"delivery_method"`
	FullName        string    `db:"full_name"`
	Email           string    `db:"email"`
	Address         string    `db:"address"`
	SubtotalPrice   int       `db:"subtotal_price"`
	TotalPrice      int       `db:"total_price"`
	DeliveryFee     int       `db:"delivery_fee"`
	Tax             int       `db:"tax"`
	Status          string    `db:"status"`
	PaymentMethod   string    `db:"payment_method"`
	ProductName     string    `db:"product_name"`
	ProductImage    string    `db:"product_image"`
	Qty             int       `db:"qty"`
	Price           int       `db:"price"`
	CreatedAt       time.Time `db:"created_at"`
}

type TransactionItem struct {
	ProductID    int    `db:"product_id"`
	ProductName  string `db:"product_name"`
	ProductImage string `db:"product_image"`
	Qty          int    `db:"qty"`
	Size         string `db:"size"`
	Variant      string `db:"variant"`
	Price        int    `db:"price"`
}
