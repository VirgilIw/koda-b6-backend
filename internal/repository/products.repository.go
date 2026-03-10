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

func (p *ProductRepository) GetProductById(ctx context.Context, id int) (model.ProductModel, error) {
	query := `
SELECT 
    id,
    name,
    description,
    price,
    created_at,
    updated_at,
    deleted_at,
    is_buy1get1,
    is_flash_sale,
    is_birthday_package
FROM products
WHERE id = $1;
`

	rows, err := p.db.Query(ctx, query, id)
	if err != nil {
		return model.ProductModel{}, err
	}
	defer rows.Close()

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ProductModel])
	if err != nil {
		return model.ProductModel{}, err
	}

	return product, nil
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

// func (p *ProductRepository) GetDetailProduct(ctx context.Context, id int) {
// query:=`select id from`
// }
