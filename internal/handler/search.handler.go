package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type SearchHandler struct {
	service *service.SearchService
}

func NewSearchHandler(service *service.SearchService) *SearchHandler {
	return &SearchHandler{
		service: service,
	}
}

// SearchProducts godoc
// @Summary Search products with filters
// @Description Get products with optional filters such as name, category, price range, and promotions
// @Tags Products
// @Produce json
// @Param name query string false "Product name"
// @Param category query string false "Product category"
// @Param min_price query int false "Minimum price"
// @Param max_price query int false "Maximum price"
// @Param is_flash_sale query bool false "Flash sale filter"
// @Param is_buy1get1 query bool false "Buy 1 Get 1 filter"
// @Param is_birthday_package query bool false "Birthday package filter"
// @Param cheap query bool false "Cheap product filter (<25000)"
// @Param page query int false "Page number"
// @Success 200 {object} dto.ResponseProductFilter
// @Failure 400 {object} dto.ResponseProductFilter
// @Failure 500 {object} dto.ResponseProductFilter
// @Router /admin/products/search [get]
func (h *SearchHandler) SearchProducts(ctx *gin.Context) {

	var req dto.SearchProductRequest
	fmt.Println("IsBirthdayPackage:", req.IsBirthdayPackage)
	fmt.Println("Raw query:", ctx.Request.URL.RawQuery)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseProductFilter{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	if req.Page < 1 {
		req.Page = 1
	}

	data, totalCount, err := h.service.SearchProducts(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseProductFilter{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	limit := 6
	totalPages := (totalCount + limit - 1) / limit

	ctx.JSON(http.StatusOK, dto.ResponseProductFilter{
		Success:    true,
		Message:    "success filter data",
		Page:       req.Page,
		TotalPages: totalPages,
		TotalCount: totalCount,
		Result:     data,
	})
}
