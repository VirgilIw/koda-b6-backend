package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

// GetCouponById godoc
// @Summary      Get Coupon
// @Description  Get Coupon By ID
// @Tags         Coupons
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Coupon ID"
// @Success      200  {object}  dto.ResponseCoupon
// @Failure      400  {object}  dto.ResponseCoupon
// @Failure      500  {object}  dto.ResponseCoupon
// @Router       /coupons/{id} [get]
func (o *OrderHandler) GetCouponById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCoupon{
			Success: false,
			Message: "bad request",
			Error:   "invalid id",
		})
		return
	}

	result, err := o.service.GetCouponById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCoupon{
			Success: false,
			Message: "Internal server error",
			Error:   "failed get coupon",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseCoupon{
		Success: true,
		Message: "success get coupon",
		Result:  result,
	})
}

// GetCoupons godoc
// @Summary      Get All Coupons
// @Description  Get list of coupons
// @Tags         Coupons
// @Produce      json
// @Success      200  {object}  dto.ResponseCoupon
// @Failure      500  {object}  dto.ResponseCoupon
// @Router       /coupons [get]
func (o *OrderHandler) GetCoupons(ctx *gin.Context) {
	result, err := o.service.GetCoupons(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCoupon{
			Success: false,
			Message: "internal server error",
			Error:   "failed get coupons",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseAllCoupon{
		Success: true,
		Message: "success get coupons",
		Result:  result,
	})
}
