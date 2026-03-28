package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{
		service: service,
	}
}

// @Summary Add product to cart
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.AddToCartRequest true "Add To Cart Request"
// @Success 200 {object} dto.ResponseCart
// @Failure 400 {object} dto.ResponseCart
// @Failure 500 {object} dto.ResponseCart
// @Router /cart [post]
func (c *CartHandler) AddToCart(ctx *gin.Context) {

	var req dto.AddToCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCart{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	userID := ctx.GetInt("userID")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, dto.ResponseCart{
			Success: false,
			Message: "unauthorized",
		})
		return
	}

	req.UserID = userID

	err := c.service.AddToCart(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCart{
			Success: false,
			Message: "failed add to cart",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseCart{
		Success: true,
		Message: "success add to cart",
	})
}

// @Summary Get user cart
// @Tags Cart
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.ResponseCart
// @Failure 401 {object} dto.ResponseCart
// @Failure 500 {object} dto.ResponseCart
// @Router /cart [get]
func (h *CartHandler) GetCart(c *gin.Context) {

	userID := c.GetInt("userID")

	carts, err := h.service.GetCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseCart{
			Success: false,
			Message: "failed to get cart",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseCart{
		Success: true,
		Message: "success get cart",
		Result:  carts,
	})
}
