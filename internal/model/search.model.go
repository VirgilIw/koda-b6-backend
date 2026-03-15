package model

type ProductFilterModel struct {
	ID                int    `db:"id"`
	Name              string `db:"name"`
	Description       string `db:"description"`
	Price             int    `db:"price"`
	Category          string `db:"category"`
	IsFlashSale       bool   `db:"is_flash_sale"`
	IsBuy1Get1        bool   `db:"is_buy1get1"`
	IsBirthdayPackage bool   `db:"is_birthday_package"`
}
