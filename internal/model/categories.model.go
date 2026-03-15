package model

import "time"

type Category struct {
	Id             int        `db:"id"`
	CategoriesName string     `db:"categories_name"`
	CreatedAt      *time.Time `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}

type CategoryExist struct {
	Exists bool `db:"exists"`
}
