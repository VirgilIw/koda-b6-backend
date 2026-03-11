package service

import (
	"context"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
)

func (p *ProductService) GetRecommendedProducts(ctx context.Context) ([]dto.ProductRecommendedResponse, error) {
	data, err := p.repo.GetRecommendedProducts(ctx)

	if err != nil {
		return nil, err
	}

	var result []dto.ProductRecommendedResponse

	for _, v := range data {
		result = append(result, dto.ProductRecommendedResponse{
			ID:        v.Id,
			Name:      v.Name,
			Price:     v.Price,
			ImagePath: v.ImagePath,
			Rating:    *v.Rating,
			Message:   *v.ReviewMessages,
		})
	}

	return result, nil
}
