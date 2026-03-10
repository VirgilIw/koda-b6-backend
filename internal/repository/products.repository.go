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

func (p *ProductRepository) GetProducts(ctx context.Context) ([]model.ProductModel, error) {
	query := `select 
    p.id, p.name, p.description, p.price,
    p.is_flash_sale, p.is_buy1get1, p.is_birthday_package, p.created_at,
    t.rating,i.image_path 
	from products p
	left join product_images pi on pi.product_id = p.id
	left join images i on i.id = pi.image_id
	left join testimonials t ON t.product_id = p.id
	order by p.id;`

	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return []model.ProductModel{}, err
	}
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.ProductModel])
	if err != nil {
		return []model.ProductModel{}, err
	}

	return products, nil
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
    string_agg(DISTINCT v.variant_name, ', ') AS variants,
    ARRAY_AGG(DISTINCT v.additional_price) AS variant_prices,
    COUNT(DISTINCT t.id) AS total_testimonials,
    ARRAY_AGG(DISTINCT s.size_name) AS sizes,
    ARRAY_AGG(DISTINCT s.additional_price) AS size_prices
FROM products p
LEFT JOIN product_variants pv ON pv.product_id = p.id
LEFT JOIN variants v ON v.id = pv.variant_id
LEFT JOIN testimonials t ON t.product_id = p.id
LEFT JOIN product_sizes ps ON ps.product_id = p.id
LEFT JOIN sizes s ON s.id = ps.size_id
WHERE p.id = $1
GROUP BY p.id, p.name,p.price;`

	rows, err := p.db.Query(ctx, query, id)

	if err != nil {
		return model.ProductDetail{}, err
	}

	detail, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ProductDetail])

	return detail, nil
}
