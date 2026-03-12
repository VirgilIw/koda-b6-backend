package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type LandingHandler struct {
	service *service.LandingService
}

func NewLandingService(service *service.LandingService) *LandingHandler {
	return &LandingHandler{
		service: service,
	}
}

// Get Recommended Product godoc
// @Summary      Get Recommended product
// @Description  Get recommended product based on minimum 3 reviews
// @Tags         LandingPage
// @Produce      json
// @Success      200  {object}  dto.ResponseRecommended
// @Failure      500  {object}  dto.ResponseRecommended
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

// GetReviews godoc
// @Summary      Get all reviews
// @Description  Retrieve all product reviews
// @Tags         LandingPage
// @Produce      json
// @Success      200  {object}  dto.ResponseReviews
// @Failure      500  {object}  dto.ResponseReviews
// @Router       /reviews [get]
func (r *LandingHandler) GetReviews(ctx *gin.Context) {

	reviews, err := r.service.GetReviews(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseReviews{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseReviews{
		Success: true,
		Message: "success get all reviews",
		Result:  reviews,
	})
}
