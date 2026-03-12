package dto

import "time"

type Size struct {
	ID              int        `json:"id,omitempty"`
	SizeName        string     `json:"size_name"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	AdditionalPrice int        `json:"additional_price"`
}

type SizeRequest struct {
	SizeName        string `json:"size_name"`
	AdditionalPrice int    `json:"additional_price"`
}

type SizeUpdateRequest struct {
	SizeName        string `json:"size_name"`
	AdditionalPrice int    `json:"additional_price"`
}
