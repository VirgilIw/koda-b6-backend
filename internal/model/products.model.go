package model

import "time"

type ProductModel struct {
	Id                int        `db:"id"`
	Name              string     `db:"name"`
	Description       string     `db:"description"`
	Price             float64    `db:"price"`
	CreatedAt         time.Time  `db:"created_at"`
	UpdatedAt         *time.Time `db:"updated_at"`
	DeletedAt         *time.Time `db:"deleted_at"`
	IsBuy1Get1        bool       `db:"is_buy1get1"`
	IsFlashSale       bool       `db:"is_flash_sale"`
	IsBirthdayPackage bool       `db:"is_birthday_package"`
}
