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

// CreateCoupon godoc
// @Summary      Create Coupon
// @Description  Create a new coupon
// @Tags         Coupons
// @Accept       json
// @Produce      json
// @Param        coupon  body      dto.CouponRequest  true  "Coupon request body"
// @Success      201     {object}  dto.ResponseCoupon
// @Failure      400     {object}  dto.ResponseCoupon
// @Failure      500     {object}  dto.ResponseCoupon
// @Router       /coupons [post]
func (o *OrderHandler) CreateCoupon(ctx *gin.Context) {
	var req dto.CouponRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCoupon{
			Success: false,
			Message: "bad request",
			Error:   "invalid request",
		})
		return
	}

	result, err := o.service.CreateCoupon(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCoupon{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponseCoupon{
		Success: true,
		Message: "success created coupon",
		Result:  result,
	})
}

// EditCoupon godoc
// @Summary      Edit Coupon
// @Description  Edit a coupon
// @Tags         Coupons
// @Accept       json
// @Produce      json
// @Param        coupon body dto.CouponRequest true "Coupon request body"
// @Success      200  {object}  dto.ResponseCoupon
// @Failure      400  {object}  dto.ResponseCoupon
// @Failure      404  {object}  dto.ResponseCoupon
// @Failure      500  {object}  dto.ResponseCoupon
// @Router       /coupons/{id} [patch]
func (o *OrderHandler) EditCoupon(ctx *gin.Context) {
	var req dto.CouponRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCoupon{
			Success: false,
			Message: "bad request",
			Error:   "invalid request",
		})
		return
	}

	result, err := o.service.EditCoupon(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCoupon{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseCoupon{
		Success: true,
		Message: "success edit coupon",
		Result:  result,
	})
}

// DeleteCoupon godoc
// @Summary      Delete Coupon
// @Description  Delete a coupon
// @Tags         Coupons
// @Accept       json
// @Produce      json
// @Param        id  	 path      int  true  "Coupon request id"
// @Success      200     {object}  dto.ResponseCoupon
// @Failure      400     {object}  dto.ResponseCoupon
// @Failure      500     {object}  dto.ResponseCoupon
// @Router       /coupons/{id} [Delete]
func (o *OrderHandler) DeleteCoupon(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCoupon{
			Success: false,
			Message: "bad request",
			Error:   "invalid id request",
		})
		return
	}

	result, err := o.service.DeleteCoupon(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCoupon{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseCoupon{
		Success: true,
		Message: "success Delete coupon",
		Result:  result,
	})
}
