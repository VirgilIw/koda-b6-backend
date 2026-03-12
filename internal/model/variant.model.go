package model

import "time"

type VariantModel struct {
	ID              int        `db:"id"`
	VariantName     string     `db:"variant_name"`
	AdditionalPrice int        `db:"additional_price"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at"`
}
