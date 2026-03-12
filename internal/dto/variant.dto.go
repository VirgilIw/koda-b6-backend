package dto

import "time"

type Variant struct {
	ID              int        `json:"id"`
	VariantName     string     `json:"variant_name"`
	AdditionalPrice int        `json:"additional_price"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}
