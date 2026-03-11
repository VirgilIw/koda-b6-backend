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
