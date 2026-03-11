package dto

import "time"

type Category struct {
	Id             int        `json:"id"`
	CategoriesName string     `json:"categories_name"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type CreateCategoryRequest struct {
	Id             int    `json:"id"`
	CategoriesName string `json:"categories_name"`
}
type CategoryRequest struct {
	CategoriesName string `json:"categories_name"`
}
type UpdateCategoryRequest struct {
	Id             int    `json:"id"`
	CategoriesName string `json:"categories_name"`
}
