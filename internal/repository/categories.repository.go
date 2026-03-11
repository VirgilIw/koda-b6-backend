package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type CategoriesRepository struct {
	db *pgxpool.Pool
}

func NewCategoriesRepository(db *pgxpool.Pool) *CategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
}

func (c *CategoriesRepository) GetCategories(ctx context.Context) ([]model.Category, error) {
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

func (c *CategoriesRepository) GetCategoryById(ctx context.Context, id int) (model.Category, error) {
	query := `SELECT 
    id,
    categories_name,
    created_at,
    updated_at,
    deleted_at
	FROM categories
	where id = $1`

	rows, err := c.db.Query(ctx, query, id)

	if err != nil {
		return model.Category{}, err
	}

	defer rows.Close()

	category, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Category])

	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (c *CategoriesRepository) CreateCategory(ctx context.Context, catName dto.CategoryRequest) (model.Category, error) {
	query := `
	INSERT INTO categories (categories_name)
	VALUES ($1)
	RETURNING 
		id,
		categories_name,
		created_at,
		updated_at,
		deleted_at
	`

	rows, err := c.db.Query(ctx, query, catName.CategoriesName)

	if err != nil {
		return model.Category{}, err
	}

	defer rows.Close()

	category, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Category])

	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (c *CategoriesRepository) UpdateCategory(ctx context.Context, req dto.CreateCategoryRequest) (model.Category, error) {
	query := `
	UPDATE categories
	SET 
		categories_name = $1,
		updated_at = now()
	WHERE id = $2
	RETURNING
		id,
		categories_name,
		created_at,
		updated_at,
		deleted_at
	`
	rows, err := c.db.Query(ctx, query, req.CategoriesName, req.Id)

	category, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Category])
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (c *CategoriesRepository) DeleteCategory(ctx context.Context, id int) error {
	query := `DELETE FROM categories
WHERE id = $1`

	_, err := c.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
