package service

import (
	"context"
	"errors"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type SizesService struct {
	repo *repository.SizesRepository
}

func NewSizesService(repo *repository.SizesRepository) *SizesService {
	return &SizesService{
		repo: repo,
	}
}

func (s *SizesService) GetSizes(ctx context.Context) ([]dto.Size, error) {
	sizes, err := s.repo.GetSizes(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.Size
	for _, size := range sizes {
		result = append(result, dto.Size{
			ID:              size.ID,
			SizeName:        size.SizeName,
			CreatedAt:       size.CreatedAt,
			UpdatedAt:       size.UpdatedAt,
			DeletedAt:       size.DeletedAt,
			AdditionalPrice: size.AdditionalPrice,
		})
	}
	return result, nil
}

func (s *SizesService) GetSizeByID(ctx context.Context, id int) (dto.Size, error) {
	if id <= 0 {
		return dto.Size{}, errors.New("invalid size id")
	}

	size, err := s.repo.GetSizeById(ctx, id)
	if err != nil {
		return dto.Size{}, err
	}

	result := dto.Size{
		ID:              size.ID,
		SizeName:        size.SizeName,
		CreatedAt:       size.CreatedAt,
		UpdatedAt:       size.UpdatedAt,
		DeletedAt:       size.DeletedAt,
		AdditionalPrice: size.AdditionalPrice,
	}

	return result, nil
}

func (s *SizesService) UpdateSize(ctx context.Context, id int, req dto.SizeUpdateRequest) (dto.Size, error) {

	if id <= 0 {
		return dto.Size{}, errors.New("invalid size id")
	}

	size, err := s.repo.UpdateSizeById(ctx, id, req)
	if err != nil {
		return dto.Size{}, err
	}

	result := dto.Size{
		ID:              size.ID,
		SizeName:        size.SizeName,
		CreatedAt:       size.CreatedAt,
		UpdatedAt:       size.UpdatedAt,
		DeletedAt:       size.DeletedAt,
		AdditionalPrice: size.AdditionalPrice,
	}

	return result, nil
}
