package service

import (
	"context"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (c *CartService) AddToCart(ctx context.Context, req dto.AddToCartRequest) error {

	err := c.repo.AddToCart(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartService) GetCart(ctx context.Context, userID int) ([]dto.CartItem, error) {

	carts, err := c.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.CartItem, 0, len(carts))

	for _, v := range carts {
		result = append(result, dto.CartItem{
			ID:           v.ID,
			UserID:       v.UserID,
			ProductID:    v.ProductID,
			ProductName:  v.ProductName,
			ProductImage: v.ProductImage,
			Price:        v.Price,
			Qty:          v.Qty,
			SizeID:       v.SizeID,
			VariantID:    v.VariantID,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return result, nil
}
