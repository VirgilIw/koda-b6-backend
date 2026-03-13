package service

import (
	"context"
	"errors"

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

func (v *VariantService) CreateVariant(ctx context.Context, req dto.VariantRequest) (dto.AllVariant, error) {
	data, err := v.repo.CreateVariant(ctx, req)

	if err != nil {
		return dto.AllVariant{}, err
	}

	var result dto.AllVariant

	result = dto.AllVariant{
		ID:              data.ID,
		VariantName:     data.VariantName,
		AdditionalPrice: data.AdditionalPrice,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
		DeletedAt:       data.DeletedAt,
	}

	return result, nil
}

func (v *VariantService) GetVariants(ctx context.Context) ([]dto.AllVariant, error) {
	variant, err := v.repo.GetVariants(ctx)

	if err != nil {
		return []dto.AllVariant{}, err
	}

	var result []dto.AllVariant

	for _, v := range variant {

		result = append(result, dto.AllVariant{
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

func (v *VariantService) GetVariantById(ctx context.Context, id int) (dto.Variant, error) {

	if id <= 0 {
		return dto.Variant{}, errors.New("id not valid")
	}

	variant, err := v.repo.GetVariantById(ctx, id)

	if err != nil {
		return dto.Variant{}, err
	}

	var result dto.Variant

	result = dto.Variant{
		ID:              variant.ID,
		VariantName:     variant.VariantName,
		AdditionalPrice: variant.AdditionalPrice,
	}

	return result, nil
}

func (v *VariantService) UpdateVariant(ctx context.Context, req dto.VariantRequest, id int) (dto.AllVariant, error) {

	if id <= 0 {
		return dto.AllVariant{}, errors.New("id not valid")
	}

	if req.VariantName == "" {
		return dto.AllVariant{}, errors.New("variant name is required")
	}

	if req.AdditionalPrice < 0 {
		return dto.AllVariant{}, errors.New("additional price cannot be under zero")
	}

	data, err := v.repo.UpdateVariant(ctx, req, id)

	if err != nil {
		return dto.AllVariant{}, errors.New("variant not found")
	}

	result := dto.AllVariant{
		ID:              data.ID,
		VariantName:     data.VariantName,
		AdditionalPrice: data.AdditionalPrice,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
		DeletedAt:       data.DeletedAt,
	}
	// fmt.Println(result)

	return result, nil
}

func (v *VariantService) DeleteVariant(ctx context.Context, id int) error {

	if id <= 0 {
		return errors.New("id not valid")
	}

	if err := v.repo.DeleteVariant(ctx, id); err != nil {
		return err
	}

	return nil
}
