package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) GetProducts(ctx context.Context, page int) ([]dto.Product, int, error) {
	data, total, err := p.repo.GetProducts(ctx, page)
	if err != nil {
		return nil, 0, err
	}

	result := make([]dto.Product, 0, len(data))

	for _, v := range data {
		var rating float64
		if v.Rating != nil {
			rating = *v.Rating
		}

		var image string
		if v.Image != nil {
			if strings.HasPrefix(*v.Image, "http") {
				image = *v.Image
			} else {
				image = "/images/" + *v.Image
			}
		}
		fmt.Println("RAW IMAGE FROM DB:", v.Image)
		sizes := v.Sizes
		if sizes == nil {
			sizes = []string{}
		}

		result = append(result, dto.Product{
			ID:                v.ID,
			Name:              v.Name,
			Rating:            rating,
			Description:       v.Description,
			Price:             v.Price,
			Stock:             v.Stock,
			IsBuy1Get1:        v.IsBuy1Get1,
			IsFlashSale:       v.IsFlashSale,
			IsBirthdayPackage: v.IsBirthdayPackage,
			Image:             image,
			Sizes:             sizes,
		})
	}

	return result, total, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, req dto.UpdateProductRequest) error {
	if err := p.repo.UpdateProduct(ctx, req); err != nil {
		return err
	}
	return nil
}

func (p *ProductService) CreateProduct(ctx context.Context, req dto.CreateProductRequest) (dto.CreateProductResponse, error) {

	product, err := p.repo.CreateProduct(ctx, req)
	if err != nil {
		return dto.CreateProductResponse{}, err
	}

	productWithSizes, err := p.repo.GetProductWithSizes(ctx, product.ID)
	if err != nil {
		return dto.CreateProductResponse{}, err
	}

	return dto.CreateProductResponse{
		ID:          productWithSizes.ID,
		Name:        productWithSizes.Name,
		Description: productWithSizes.Description,
		Price:       productWithSizes.Price,
		Sizes:       productWithSizes.Sizes,
		Images:      req.Images,
		Stock:       productWithSizes.Stock,
	}, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, id int) (err error) {
	tx, err := p.repo.Begin(ctx)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	if err = p.repo.DeleteProductCategoriesTx(ctx, tx, id); err != nil {
		return
	}

	if err = p.repo.DeleteProductSizesTx(ctx, tx, id); err != nil {
		return
	}

	if err = p.repo.DeleteProductVariantsTx(ctx, tx, id); err != nil {
		return
	}

	if err = p.repo.DeleteProductImagesTx(ctx, tx, id); err != nil {
		return
	}

	if err = p.repo.DeleteProductTestimonialsTx(ctx, tx, id); err != nil {
		return
	}

	if err = p.repo.DeleteProductTx(ctx, tx, id); err != nil {
		return
	}

	return
}

func mapSizes(names []string, prices []int) []dto.SizeResponse {
	var result []dto.SizeResponse

	for i := range names {
		price := 0
		if i < len(prices) {
			price = prices[i]
		}

		result = append(result, dto.SizeResponse{
			Name:  names[i],
			Price: price,
		})
	}

	return result
}

func mapVariants(names []string, prices []int) []dto.VariantResponse {
	var result []dto.VariantResponse

	for i := range names {
		result = append(result, dto.VariantResponse{
			Name:  names[i],
			Price: prices[i],
		})
	}

	return result
}

func (p *ProductService) GetDetailProductById(ctx context.Context, id int, selectedSize string, selectedVariant string) (dto.ProductDetail, error) {
	detail, err := p.repo.GetDetailProductById(ctx, id)
	if err != nil {
		return dto.ProductDetail{}, err
	}

	productDetail := dto.ProductDetail{
		ID:           detail.ID,
		Name:         detail.Name,
		Images:       detail.Images,
		Description:  detail.Description,
		Rating:       detail.Rating,
		Price:        detail.Price,
		TotalReviews: detail.TotalReviews,

		Sizes:    mapSizes(detail.Sizes, detail.SizePrices),
		Variants: mapVariants(detail.Variants, detail.VariantPrices),
	}

	return productDetail, nil
}
