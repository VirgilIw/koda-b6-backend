package service

import (
	"context"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type VariantService struct {
	repo *repository.VariantRepository
}

func NewVariantService(repo *repository.VariantRepository) *VariantService {
	return &VariantService{
		repo: repo,
	}
}

func (v *VariantService) GetVariants(ctx context.Context) ([]dto.Variant, error) {
	variant, err := v.repo.GetVariants(ctx)

	if err != nil {
		return []dto.Variant{}, err
	}

	var result []dto.Variant

	for _, v := range variant {

		result = append(result, dto.Variant{
			ID:              v.ID,
			VariantName:     v.VariantName,
			AdditionalPrice: v.AdditionalPrice,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
			DeletedAt:       v.DeletedAt,
		})
	}

	return result, nil
}
