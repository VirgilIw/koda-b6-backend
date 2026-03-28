package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductService(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// GetProducts godoc
// @Summary      Get products
// @Description  Get products with pagination
// @Tags         Products
// @Produce      json
// @Param        page     query     int     false  "Page number"
// @Success      200      {object}  dto.ProductsResponse
// @Failure      400      {object}  dto.ProductsResponse
// @Failure      500      {object}  dto.ProductsResponse
// @Security     BearerAuth
// @Router       /admin/products [get]
func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	pageStr := ctx.Query("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	products, err := p.service.GetProducts(ctx.Request.Context(), page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ProductsResponse{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ProductsResponse{
		Success: true,
		Message: "Success get data",
		Result:  products,
	})
}

// bad request dari user
// internal server dari kode kita

// GetProductsById godoc
// @Summary      Get product by ID size variant
// @Description  Get a single product by ID and calculate price based on selected size and variant
// @Tags         Products
// @Produce      json
// @Param        id       path   int     true  "Product ID"
// @Param        size     query  string false "Size" example(Regular/Medium/Large)
// @Param        variant  query  string false "Variant" example(Hot/Iced)
// @Success      200  {object}  dto.ProductDetailResponse
// @Failure      400  {object}  dto.ProductsResponse
// @Failure      404  {object}  dto.ProductsResponse
// @Failure      500  {object}  dto.ProductsResponse
// @Security     BearerAuth
// @Router       /admin/products/{id} [get]
func (p *ProductHandler) GetDetailProductById(ctx *gin.Context) {
	// ambil id dari path
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, dto.ProductsResponse{
			Success: false,
			Message: "invalid id",
			Result:  nil,
		})
		return
	}

	// ambil selectedSize dan selectedVariant dari query params
	selectedSize := ctx.Query("size")
	selectedVariant := ctx.Query("variant")

	product, err := p.service.GetDetailProductById(ctx, id, selectedSize, selectedVariant)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ProductDetailResponse{
			Success: false,
			Message: "product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ProductDetailResponse{
		Success: true,
		Message: "success get data by id",
		Result:  product,
	})
}

// UpdateProduct godoc
// @Summary      Update product
// @Description  Update existing product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id      path  int  true  "Product ID"
// @Param        request body  dto.UpdateProductRequest true "Update product request"
// @Success      200  {object}  dto.ProductsResponse
// @Failure      400  {object}  dto.ProductsResponse
// @Failure      500  {object}  dto.ProductsResponse
// @Security     BearerAuth
// @Router       /admin/products/{id} [patch]
func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var req dto.UpdateProductRequest

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ProductsResponse{
			Success: false,
			Message: "bad request",
			Error:   "invalid product id",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ProductsResponse{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	req.Id = id

	if err := p.service.UpdateProduct(ctx, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ProductsResponse{
			Success: false,
			Message: "failed update product",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ProductsResponse{
		Success: true,
		Message: "success update product",
	})
}

// UpdateProduct godoc
// @Summary      Create product
// @Description  Create new product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request body  dto.CreateProductRequest true "Create product request"
// @Success      201  {object}  dto.SingleProductResponse
// @Failure      400  {object}  dto.SingleProductResponse
// @Failure      500  {object}  dto.SingleProductResponse
// @Security     BearerAuth
// @Router       /admin/products [post]
func (p *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ProductsResponse{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	data, err := p.service.CreateProduct(ctx.Request.Context(), req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ProductsResponse{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.SingleProductResponse{
		Success: true,
		Message: "success create product",
		Result:  data,
	})
}

// DeleteProduct godoc
// @Summary      Delete product
// @Description  Delete product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id      path  int  true  "Product ID"
// @Success      200  {object}  dto.ProductsResponse
// @Failure      400  {object}  dto.ProductsResponse
// @Failure      500  {object}  dto.ProductsResponse
// @Security     BearerAuth
// @Router       /admin/products [delete]
func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ProductsResponse{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	if err := p.service.DeleteProduct(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ProductsResponse{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ProductsResponse{
		Success: true,
		Message: "delete product success",
	})
}
