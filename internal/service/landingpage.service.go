package service

import (
	"context"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type LandingService struct {
	repo *repository.LandingRepository
}

func NewLandingService(repo *repository.LandingRepository) *LandingService {
	return &LandingService{
		repo: repo,
	}
}

func (p *ProductService) GetRecommendedProducts(ctx context.Context) ([]dto.ProductRecommendedResponse, error) {
	data, err := p.repo.GetRecommendedProducts(ctx)

	if err != nil {
		return []dto.ProductRecommendedResponse{}, err
	}

	var result = []dto.ProductRecommendedResponse{}

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

func (r *LandingService) GetReviews(ctx context.Context) ([]dto.Reviews, error) {
	lanRepo, err := r.repo.GetReviews(ctx)

	if err != nil {
		return []dto.Reviews{}, err
	}

	result := []dto.Reviews{}

	for _, v := range lanRepo {
		result = append(result, dto.Reviews{
			ID:          v.ID,
			Name:        v.Name,
			Image:       v.Image,
			AuthorTitle: v.AuthorTitle,
			Message:     v.Message,
			Rating:      v.Rating,
			CreatedAt:   v.CreatedAt,
		})
	}

	return result, nil
}
