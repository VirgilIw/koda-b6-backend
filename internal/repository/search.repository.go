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

func (r *SearchRepository) SearchProducts(ctx context.Context, req dto.SearchProductRequest) ([]model.ProductFilterModel, int, error) {

	whereClause := " WHERE 1=1"
	args := []any{}
	i := 1

	// Normalize (defensive)
	if req.Page < 1 {
		req.Page = 1
	}

	// Dynamic Filters

	if req.Name != "" {
		whereClause += fmt.Sprintf(" AND p.name ILIKE $%d", i)
		args = append(args, "%"+req.Name+"%")
		i++
	}

	if req.Category != "" {
		whereClause += fmt.Sprintf(` AND p.id IN (
			SELECT pc2.product_id FROM product_categories pc2
			JOIN categories c2 ON c2.id = pc2.categories_id
			WHERE c2.categories_name ILIKE $%d
		)`, i)
		args = append(args, "%"+req.Category+"%")
		i++
	}

	if req.MinPrice > 0 {
		whereClause += fmt.Sprintf(" AND p.price >= $%d", i)
		args = append(args, req.MinPrice)
		i++
	}

	if req.MaxPrice > 0 {
		whereClause += fmt.Sprintf(" AND p.price <= $%d", i)
		args = append(args, req.MaxPrice)
		i++
	}

	if req.IsFlashSale {
		whereClause += " AND p.is_flash_sale = true"
	}

	if req.IsBuy1Get1 {
		whereClause += " AND p.is_buy1get1 = true"
	}

	if req.Recommended {
		whereClause += `
    AND p.id IN (
        SELECT t2.product_id
        FROM testimonials t2
        GROUP BY t2.product_id
        HAVING AVG(t2.rating) >= 4.6
    )`
	}

	if req.IsBirthdayPackage {
		whereClause += " AND p.is_birthday_package = true"
	}

	if req.Cheap {
		whereClause += " AND p.price <= 20000"
	}

	// COUNT QUERY
	countQuery := `
SELECT COUNT(DISTINCT p.id)
FROM products p
LEFT JOIN product_categories pc ON pc.product_id = p.id
LEFT JOIN categories c ON c.id = pc.categories_id
` + whereClause

	countRows, err := r.db.Query(ctx, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer countRows.Close()

	totalCount, err := pgx.CollectOneRow(countRows, pgx.RowTo[int])
	if err != nil {
		return nil, 0, err
	}

	// PAGINATION

	limit := 6
	offset := (req.Page - 1) * limit

	// DATA QUERY

	dataQuery := `
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
` + whereClause +
		fmt.Sprintf(`
GROUP BY 
    p.id
ORDER BY p.id ASC
LIMIT 6 OFFSET $%d
`, i)

	dataArgs := append(args, offset)

	rows, err := r.db.Query(ctx, dataQuery, dataArgs...)
	if err != nil {
		return nil, 0, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.ProductFilterModel])
	if err != nil {
		return nil, 0, err
	}

	for i := range products {
		if products[i].Categories == nil {
			products[i].Categories = []string{}
		}
		if products[i].Images == nil {
			products[i].Images = []string{}
		}
	}

	return products, totalCount, nil
}
