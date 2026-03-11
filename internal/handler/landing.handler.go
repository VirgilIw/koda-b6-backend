package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
)

// Get Recommended Product godoc
// @Summary      Get Recommended product
// @Description  Get recommended product based on minimum 3 reviews
// @Tags         Recommended Products
// @Produce      json
// @Success      200  {object}  dto.ResponseRecommended
// @Failure      500  {object}  dto.ResponseRecommended
// @Security     BearerAuth
// @Router       /recommended-products [get]
func (h *ProductHandler) GetRecommendedProducts(ctx *gin.Context) {

	products, err := h.service.GetRecommendedProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseRecommended{
			Success: false,
			Message: "failed to get recommended products",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseRecommended{
		Success: true,
		Message: "success get recommended products",
		Result:  products,
	})
}
