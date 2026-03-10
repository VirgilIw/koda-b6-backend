package service

import (
	"context"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (o *OrderService) GetCouponById(ctx context.Context, id int) (dto.Coupon, error) {
	coupon, err := o.repo.GetCouponById(ctx, id)

	if err != nil {
		return dto.Coupon{}, err
	}

	var data dto.Coupon

	data = dto.Coupon{
		ID:          coupon.ID,
		Title:       coupon.Title,
		Description: coupon.Description,
		Value:       coupon.Value,
		Image:       coupon.Image,
		CreatedAt:   coupon.CreatedAt,
	}

	return data, nil
}
