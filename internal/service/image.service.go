package service

import (
	"context"
	"errors"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type ImageService struct {
	repo *repository.ImageRepository
}

func NewImageService(repo *repository.ImageRepository) *ImageService {
	return &ImageService{
		repo: repo,
	}
}

func (i *ImageService) CreateImage(ctx context.Context, req dto.ImageRequest) (dto.ImageDto, error) {
	data, err := i.repo.CreateImage(ctx, req)

	if err != nil {
		return dto.ImageDto{}, err
	}

	var result dto.ImageDto

	result = dto.ImageDto{
		Id:        data.Id,
		ImagePath: data.ImagePath,
	}

	return result, nil
}

func (i *ImageService) GetImages(ctx context.Context) ([]dto.ImageDto, error) {
	data, err := i.repo.GetImages(ctx)

	if err != nil {
		return []dto.ImageDto{}, err
	}

	var result []dto.ImageDto

	for _, v := range data {
		result = append(result, dto.ImageDto{
			Id:        v.Id,
			ImagePath: v.ImagePath,
		})
	}

	return result, nil
}

func (i *ImageService) GetImageById(ctx context.Context, id int) (dto.ImageDto, error) {
	if id <= 0 {
		return dto.ImageDto{}, errors.New("id not found")
	}

	data, err := i.repo.GetImageById(ctx, id)
	if err != nil {
		return dto.ImageDto{}, err
	}

	var result dto.ImageDto

	result = dto.ImageDto{
		Id:        data.Id,
		ImagePath: data.ImagePath,
	}

	return result, nil

}

func (i *ImageService) UpdateImage(ctx context.Context, req dto.ImageRequest, id int) (dto.ImageDto, error) {

	if req.ImagePath == "" {
		return dto.ImageDto{}, errors.New("image path is required")
	}

	data, err := i.repo.UpdateImage(ctx, req, id)
	if err != nil {
		return dto.ImageDto{}, err
	}

	return dto.ImageDto{
		Id:        data.Id,
		ImagePath: data.ImagePath,
	}, nil
}

func (i *ImageService) DeleteImageById(ctx context.Context, id int) error {

	if id <= 0 {
		return errors.New("id not found")
	}

	if err := i.repo.DeleteImageById(ctx, id); err != nil {
		return err
	}

	return nil
}
