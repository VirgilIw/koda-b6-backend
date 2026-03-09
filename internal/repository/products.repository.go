package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
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
	query := `SELECT 
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
	FROM products Order by id;`
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
