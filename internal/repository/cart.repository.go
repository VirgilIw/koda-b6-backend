package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (c *CartRepository) AddToCart(ctx context.Context, req dto.AddToCartRequest) error {

	query := `
	INSERT INTO cart (user_id, product_id, qty, size_id, variant_id)
	VALUES ($1,$2,$3,$4,$5)
	ON CONFLICT (user_id, product_id, size_id, variant_id)
	DO UPDATE SET qty = cart.qty + EXCLUDED.qty
	`

	_, err := c.db.Exec(ctx, query,
		req.UserID,
		req.ProductID,
		req.Qty,
		req.SizeID,
		req.VariantID,
	)

	return err
}

func (c *CartRepository) GetCartByUserID(ctx context.Context, userID int) ([]model.CartItem, error) {

	query :=
		`
		SELECT
		c.id,
		c.user_id,
		c.product_id,
		p.name AS name,
		i.image_path AS image_path,
		p.price,
		c.qty,
		c.size_id,
		c.variant_id,
		c.created_at,
		c.updated_at,
		c.deleted_at
		FROM cart c
		JOIN products p ON p.id = c.product_id
		LEFT JOIN product_images pi ON pi.product_id = p.id
		LEFT JOIN images i ON i.id = pi.image_id
		WHERE c.user_id = $1
	`

	rows, err := c.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.CartItem])
	if err != nil {
		return nil, err
	}

	return data, nil
}
