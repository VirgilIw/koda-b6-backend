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

func (p *ProductRepository) GetProducts(ctx context.Context, page int) ([]model.ProductModel, int, error) {

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
    ARRAY_AGG(DISTINCT s.size_name) FILTER (WHERE s.size_name IS NOT NULL) AS sizes,
    p.is_birthday_package,
    AVG(t.rating) AS rating,
	COUNT(p.id) OVER() AS total,
	p.stock AS stock,
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
LEFT JOIN product_sizes ps ON ps.product_id = p.id
LEFT JOIN sizes s ON s.id = ps.size_id
GROUP BY 
    p.id, p.name, p.description, p.price, p.created_at,
    p.is_flash_sale, p.is_buy1get1, p.is_birthday_package
ORDER BY p.id
LIMIT 6 OFFSET $1;
	`

	rows, err := p.db.Query(ctx, query, offset)
	if err != nil {
		return nil, 0, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.ProductModel])
	if err != nil {
		return nil, 0, err
	}

	total := 0
	if len(products) > 0 {
		total = products[0].Total
	}

	return products, total, nil
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

func (p *ProductRepository) CreateProduct(ctx context.Context, req dto.CreateProductRequest) (model.CreateProductModel, error) {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return model.CreateProductModel{}, err
	}
	defer tx.Rollback(ctx)

	// 1. Insert product
	productQuery := `
	INSERT INTO products (
		name,
		description,
		price,
		stock
	) VALUES ($1, $2, $3, $4)
	RETURNING id, name, description, price, stock, created_at
	`

	rows, err := tx.Query(
		ctx,
		productQuery,
		req.Name,
		req.Description,
		req.Price,
		req.Stock, // ✅ FIX DI SINI
	)
	if err != nil {
		return model.CreateProductModel{}, err
	}

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.CreateProductModel])
	if err != nil {
		return model.CreateProductModel{}, err
	}

	// 2. Insert sizes
	if len(req.Sizes) > 0 {
		for _, sizeID := range req.Sizes {
			_, err := tx.Exec(
				ctx,
				`
				INSERT INTO product_sizes (product_id, size_id)
				VALUES ($1, $2)
				`,
				product.ID,
				sizeID,
			)
			if err != nil {
				return model.CreateProductModel{}, err
			}
		}
	}

	// 3. Handle image (optional)
	if req.Images != nil {
		imageRows, err := tx.Query(
			ctx,
			`
			INSERT INTO images (image_path)
			VALUES ($1)
			RETURNING id
			`,
			*req.Images, // filename
		)
		if err != nil {
			return model.CreateProductModel{}, err
		}

		imageID, err := pgx.CollectOneRow(imageRows, pgx.RowTo[int64])
		if err != nil {
			return model.CreateProductModel{}, err
		}

		// insert relation
		_, err = tx.Exec(
			ctx,
			`
			INSERT INTO product_images (product_id, image_id)
			VALUES ($1, $2)
			`,
			product.ID,
			imageID,
		)
		if err != nil {
			return model.CreateProductModel{}, err
		}
	}

	// 4. Commit
	if err := tx.Commit(ctx); err != nil {
		return model.CreateProductModel{}, err
	}

	return product, nil
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
