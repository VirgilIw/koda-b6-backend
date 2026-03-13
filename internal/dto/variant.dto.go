package dto

import "time"

type AllVariant struct {
	ID              int        `json:"id,omitempty"`
	VariantName     string     `json:"variant_name,omitempty"`
	AdditionalPrice int        `json:"additional_price,omitempty"`
	CreatedAt       time.Time  `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

type Variant struct {
	ID              int    `json:"id,omitempty"`
	VariantName     string `json:"variant_name,omitempty"`
	AdditionalPrice int    `json:"additional_price,omitempty"`
}

type VariantRequest struct {
	VariantName     string `json:"variant_name,omitempty"`
	AdditionalPrice int    `json:"additional_price,omitempty"`
}
