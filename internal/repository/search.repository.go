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
SELECT 
    p.id, 
    p.name, 
    p.description, 
    p.price,
    p.is_flash_sale, 
    p.is_buy1get1, 
    p.is_birthday_package,

    AVG(t.rating) AS rating,

    ARRAY_AGG(DISTINCT c.categories_name ORDER BY c.categories_name)
        FILTER (WHERE c.id IS NOT NULL) AS categories,

    ARRAY_AGG(DISTINCT i.image_path ORDER BY i.image_path)
        FILTER (WHERE i.id IS NOT NULL) AS images

FROM products p

LEFT JOIN testimonials t ON t.product_id = p.id
LEFT JOIN product_categories pc ON pc.product_id = p.id
LEFT JOIN categories c ON c.id = pc.categories_id
LEFT JOIN product_images pi ON pi.product_id = p.id
LEFT JOIN images i ON i.id = pi.image_id

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
		query += fmt.Sprintf(` AND p.id IN (
            SELECT pc2.product_id FROM product_categories pc2
            JOIN categories c2 ON c2.id = pc2.categories_id
            WHERE c2.categories_name ILIKE $%d
        )`, i)
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

	query += " GROUP BY p.id ORDER BY p.name ASC"

	var limit int
	limit = 6

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

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.ProductFilterModel])
}
