package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type CategorieRepository struct {
	db *pgxpool.Pool
}

func NewCategoriesRepository(db *pgxpool.Pool) *CategorieRepository {
	return &CategorieRepository{
		db: db,
	}
}

func (c *CategorieRepository) GetCategories(ctx context.Context) ([]model.Category, error) {
	query := `SELECT 
    id,
    categories_name,
    created_at,
    updated_at,
    deleted_at
FROM categories;`

	rows, err := c.db.Query(ctx, query)

	if err != nil {
		return []model.Category{}, err
	}

	category, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Category])

	if err != nil {
		return []model.Category{}, err
	}

	return category, nil
}
