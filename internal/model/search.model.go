package model

type ProductFilterModel struct {
	ID                int      `db:"id"`
	Name              string   `db:"name"`
	Description       string   `db:"description"`
	Price             int      `db:"price"`
	Categories        []string `db:"categories"`
	Images            []string `db:"images"`
	Rating            *float64 `db:"rating"`
	IsFlashSale       bool     `db:"is_flash_sale"`
	IsBuy1Get1        bool     `db:"is_buy1get1"`
	IsBirthdayPackage bool     `db:"is_birthday_package"`
}
