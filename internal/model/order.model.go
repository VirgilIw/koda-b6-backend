package model

import "time"

type Coupon struct {
	ID          int        `db:"id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Value       string     `db:"value"`
	Image       string     `db:"image"`
	CreatedAt   *time.Time `db:"created_at"`
}
