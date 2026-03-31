package model

import "time"

type TransactionProduct struct {
	ID            int       `db:"id"`
	ProductID     int       `db:"product_id"`
	TransactionID int       `db:"transaction_id"`
	Qty           int       `db:"qty"`
	Size          string    `db:"size"`
	Variant       string    `db:"variant"`
	Price         int       `db:"price"`
	CreatedAt     time.Time `db:"created_at"`
}
