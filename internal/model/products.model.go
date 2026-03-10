package model

import "time"

type ProductModel struct {
	Id                int       `db:"id"`
	Name              string    `db:"name"`
	Description       string    `db:"description"`
	Price             float64   `db:"price"`
	IsFlashSale       bool      `db:"is_flash_sale"`
	IsBuy1Get1        bool      `db:"is_buy1get1"`
	IsBirthdayPackage bool      `db:"is_birthday_package"`
	CreatedAt         time.Time `db:"created_at"`
	ImagePath         *string   `db:"image_path"`
	Rating            *float64  `db:"rating"`
}
