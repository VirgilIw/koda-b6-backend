package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type SearchRepository struct {
	db *pgxpool.Pool
}

func NewSearchRepository(db *pgxpool.Pool) *SearchRepository {
	return &SearchRepository{
		db: db,
	}
}

func (r *SearchRepository) SearchProducts(ctx context.Context, req dto.SearchProductRequest) ([]model.ProductFilterModel, error) {

	query := `
SELECT DISTINCT ON (p.id)
	p.id,
	p.name,
	p.description,
	p.price,
	c.categories_name AS category,
	p.is_flash_sale,
	p.is_buy1get1,
	p.is_birthday_package
FROM products p
LEFT JOIN product_categories pc ON pc.product_id = p.id
LEFT JOIN categories c ON c.id = pc.categories_id
WHERE 1=1
`

	args := []any{}
	i := 1

	if req.Name != "" {
		query += fmt.Sprintf(" AND p.name ILIKE $%d", i)
		args = append(args, "%"+req.Name+"%")
		i++
	}

	if req.Category != "" {
		query += fmt.Sprintf(" AND c.categories_name ILIKE $%d", i)
		args = append(args, "%"+req.Category+"%")
		i++
	}

	if req.MinPrice > 0 {
		query += fmt.Sprintf(" AND p.price >= $%d", i)
		args = append(args, req.MinPrice)
		i++
	}

	if req.MaxPrice > 0 {
		query += fmt.Sprintf(" AND p.price <= $%d", i)
		args = append(args, req.MaxPrice)
		i++
	}

	if req.IsFlashSale {
		query += " AND p.is_flash_sale = true"
	}

	if req.IsBuy1Get1 {
		query += " AND p.is_buy1get1 = true"
	}

	if req.IsBirthdayPackage {
		query += " AND p.is_birthday_package = true"
	}

	if req.Cheap {
		query += " AND p.price < 25000"
	}

	// PostgreSQL rule: DISTINCT ON column must appear first in ORDER BY
	query += " ORDER BY p.id, p.name ASC"

	limit := req.Limit
	if limit == 0 {
		limit = 4
	}

	page := req.Page
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * limit

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)

	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.ProductFilterModel])
	if err != nil {
		return nil, err
	}

	return products, nil
}
