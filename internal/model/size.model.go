package model

import "time"

type Size struct {
	ID              int        `db:"id"`
	SizeName        string     `db:"size_name"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at"`
	AdditionalPrice int        `db:"additional_price"`
}

type ProductWithSizes struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int       `db:"price"`
	Stock       int       `db:"stock"`
	CreatedAt   time.Time `db:"created_at"`
	Sizes       []string  `db:"sizes"`
}
