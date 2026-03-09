package handler

import (
	"net/http"

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
// @Description  Get products
// @Tags         Products
// @Produce      json
// @Success      200  {object}  dto.ProductsResponse
// @Failure      400  {object}  dto.ProductsResponse
// @Failure      500  {object}  dto.ProductsResponse
// @Security     BearerAuth
// @Router       /products [get]
func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := p.service.GetProducts(ctx.Request.Context())
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
