package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type ProductRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewProductRepository(db *pgxpool.Pool, rdb *redis.Client) *ProductRepository {
	return &ProductRepository{
		db:  db,
		rdb: rdb,
	}
}
func (p *ProductRepository) GetProducts(ctx context.Context, page int) ([]model.ProductModel, error) {

	if page == 0 {
		page = 1
	}

	limit := 6
	offset := (page - 1) * limit

	query := `
SELECT 
    p.id,
    p.name,
    p.description,
    p.price,
    p.created_at,
    p.is_flash_sale,
    p.is_buy1get1,
    p.is_birthday_package,
    AVG(t.rating) AS rating,
    (
        SELECT i.image_path
        FROM product_images pi
        JOIN images i ON i.id = pi.image_id
        WHERE pi.product_id = p.id
        ORDER BY i.id
        LIMIT 1
    ) AS image
FROM products p
LEFT JOIN testimonials t ON t.product_id = p.id
GROUP BY 
    p.id, p.name, p.description, p.price, p.created_at,
    p.is_flash_sale, p.is_buy1get1, p.is_birthday_package

ORDER BY p.id
LIMIT 6 OFFSET $1;
	`

	rows, err := p.db.Query(ctx, query, offset)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.ProductModel])
}

func (p *ProductRepository) UpdateProduct(ctx context.Context, req dto.UpdateProductRequest) error {
	query := `
	UPDATE products
	SET
		name = $1,
		description = $2,
		price = $3,
		updated_at = NOW(),
		is_buy1get1 = $4,
		is_flash_sale = $5,
		is_birthday_package = $6
	WHERE id = $7
	`

	cmdTag, err := p.db.Exec(
		ctx,
		query,
		req.Name,
		req.Description,
		req.Price,
		req.IsBuy1Get1,
		req.IsFlashSale,
		req.IsBirthdayPackage,
		req.Id,
	)

	if err != nil {
		return err
	}

	// cek apakah ada row yang diupdate
	if cmdTag.RowsAffected() == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (p *ProductRepository) CreateProduct(ctx context.Context, req dto.CreateProductRequest) (model.ProductModel, error) {
	query := `
INSERT INTO products (
    name,
    description,
    price,
    is_buyget1,
    is_flash_sale,
    is_birthday_package
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, description, price, is_buyget1, is_flash_sale, is_birthday_package, created_at
`

	rows, err := p.db.Query(
		ctx,
		query,
		req.Name,
		req.Description,
		req.Price,
		req.IsBuy1Get1,
		req.IsFlashSale,
		req.IsBirthdayPackage,
	)
	if err != nil {
		return model.ProductModel{}, err
	}
	defer rows.Close()

	productModel, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ProductModel])
	if err != nil {
		return model.ProductModel{}, err
	}

	return productModel, nil
}

func (p *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM PRODUCTS WHERE ID=$1`
	_, err := p.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) GetDetailProductById(ctx context.Context, id int) (model.ProductDetail, error) {
	query := `
SELECT 
    p.id,
    p.name,
    p.price,
	p.description,
    ARRAY_AGG(DISTINCT v.variant_name) AS variants,
    ARRAY_AGG(DISTINCT v.additional_price) AS variant_prices,
    COUNT(DISTINCT t.id) AS total_testimonials,
    ARRAY_AGG(DISTINCT s.size_name) AS sizes,
    ARRAY_AGG(DISTINCT s.additional_price) AS size_prices,
    (
        SELECT ARRAY_AGG(i.image_path ORDER BY i.id)
        FROM product_images pi
        JOIN images i ON i.id = pi.image_id
        WHERE pi.product_id = p.id
    ) AS images
FROM products p
LEFT JOIN product_variants pv ON pv.product_id = p.id
LEFT JOIN variants v ON v.id = pv.variant_id
LEFT JOIN testimonials t ON t.product_id = p.id
LEFT JOIN product_sizes ps ON ps.product_id = p.id
LEFT JOIN sizes s ON s.id = ps.size_id
WHERE p.id = $1
GROUP BY p.id, p.name, p.price, p.description;`

	rows, err := p.db.Query(ctx, query, id)

	if err != nil {
		return model.ProductDetail{}, err
	}

	detail, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ProductDetail])

	return detail, nil
}

func (p *ProductRepository) GetRecommendedProducts(ctx context.Context) ([]model.RecommendedProductModel, error) {
	query := `SELECT 
    p.id,
    p.name,
    p.description,
    p.price,
    AVG(t.rating) AS rating,
    STRING_AGG(t.message, ' , ') AS review_messages,
    (
        SELECT i.image_path
        FROM product_images pi
        JOIN images i ON i.id = pi.image_id
        WHERE pi.product_id = p.id
        LIMIT 1
    ) AS image_path,
    COUNT(DISTINCT t.id) AS total_review
FROM products p
LEFT JOIN testimonials t 
    ON t.product_id = p.id
GROUP BY p.id, p.name, p.description, p.price
HAVING COUNT(DISTINCT t.id) >= 3
ORDER BY total_review DESC
LIMIT 4;`

	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return []model.RecommendedProductModel{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.RecommendedProductModel])

	if err != nil {
		return []model.RecommendedProductModel{}, err
	}

	if products == nil {
		products = []model.RecommendedProductModel{}
	}

	return products, nil
}
