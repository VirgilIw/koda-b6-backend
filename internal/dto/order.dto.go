package dto

import "time"

type Coupon struct {
	ID          int        `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Value       string     `json:"value,omitempty"`
	Image       string     `json:"image,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}

type CouponRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       string `json:"value"`
	Image       string `json:"image"`
}
