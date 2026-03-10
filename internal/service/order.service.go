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

func (o *OrderService) GetCoupons(ctx context.Context) ([]dto.Coupon, error) {
	coupon, err := o.repo.GetCoupons(ctx)

	if err != nil {
		return []dto.Coupon{}, err
	}

	var data []dto.Coupon

	for _, v := range coupon {
		data = append(data, dto.Coupon{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Value:       v.Value,
			Image:       v.Image,
			CreatedAt:   v.CreatedAt,
		})
	}

	return data, nil
}

func (o *OrderService) CreateCoupon(ctx context.Context, req dto.CouponRequest) (dto.Coupon, error) {
	coupon, err := o.repo.CreateCoupon(ctx, req)

	if err != nil {
		return dto.Coupon{}, err
	}

	var result dto.Coupon

	result = dto.Coupon{
		ID:          coupon.ID,
		Title:       coupon.Title,
		Description: coupon.Description,
		Value:       coupon.Value,
		Image:       coupon.Image,
	}

	return result, nil
}

func (o *OrderService) EditCoupon(ctx context.Context, req dto.CouponRequest) (dto.Coupon, error) {
	coupon, err := o.repo.EditCoupon(ctx, req)

	if err != nil {
		return dto.Coupon{}, err
	}

	var result dto.Coupon

	result = dto.Coupon{
		ID:          coupon.ID,
		Title:       coupon.Title,
		Description: coupon.Description,
		Value:       coupon.Value,
		Image:       coupon.Image,
	}

	return result, nil
}

func (o *OrderService) DeleteCoupon(ctx context.Context, id int) (dto.Coupon, error) {
	coupon, err := o.repo.DeleteCoupon(ctx, id)

	if err != nil {
		return dto.Coupon{}, err
	}

	var result dto.Coupon

	result = dto.Coupon{
		ID:          coupon.ID,
		Title:       coupon.Title,
		Description: coupon.Description,
		Value:       coupon.Value,
		Image:       coupon.Image,
	}

	return result, nil
}
