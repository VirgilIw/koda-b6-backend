package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type VariantHandler struct {
	service *service.VariantService
}

func NewVariantHandler(service *service.VariantService) *VariantHandler {
	return &VariantHandler{
		service: service,
	}
}

// Create New Variant godoc
// @Summary      Create new variant
// @Description  Create new variant
// @Accept       json
// @Tags         Variants
// @Produce      json
// @Param        req body dto.VariantRequest true "Create variant request"
// @Success      201  {object}  dto.ResponseUpdateVariant
// @Failure      400  {object}  dto.ResponseUpdateVariant
// @Failure      500  {object}  dto.ResponseUpdateVariant
// @Security     BearerAuth
// @Router       /admin/variants [post]
func (v *VariantHandler) CreateVariant(ctx *gin.Context) {
	var req dto.VariantRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseCreateVariant{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	data, err := v.service.CreateVariant(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseCreateVariant{
			Success: false,
			Message: "internal server error",
			Error:   "failed to create variant",
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponseCreateVariant{
		Success: true,
		Message: "success create variant",
		Result:  data,
	})
}

// Get All Variants godoc
// @Summary      Get all variants
// @Description  Retrieve all variants
// @Tags         Variants
// @Produce      json
// @Success      200  {object}  dto.ResponseVariants
// @Failure      500  {object}  dto.ResponseVariants
// @Security     BearerAuth
// @Router       /admin/variants [get]
func (v *VariantHandler) GetVariants(ctx *gin.Context) {

	variants, err := v.service.GetVariants(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseVariants{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseVariants{
		Success: true,
		Message: "success get all variants",
		Result:  variants,
	})

}

// Get Variant By Id godoc
// @Summary      Get variant by id
// @Description  Retrieve variant by id
// @Tags         Variants
// @Produce      json
// @Param        id   path  int  true  "variant id"
// @Success      200  {object}  dto.ResponseVariant
// @Failure      400  {object}  dto.ResponseVariant
// @Failure      500  {object}  dto.ResponseVariant
// @Security     BearerAuth
// @Router       /admin/variants/{id} [get]
func (v *VariantHandler) GetVariantById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseVariant{
			Success: false,
			Message: "invalid id",
			Error:   err.Error(),
		})
		return
	}

	variant, err := v.service.GetVariantById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseVariant{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseVariant{
		Success: true,
		Message: "success get variant by id",
		Result:  variant,
	})
}

// Update Variant By Id godoc
// @Summary      Update variant by id
// @Description  Update variant by id
// @Tags         Variants
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "variant id"
// @Param        req  body  dto.VariantRequest true "update variant request"
// @Success      200  {object}  dto.ResponseCreateVariant
// @Failure      400  {object}  dto.ResponseCreateVariant
// @Failure      404  {object}  dto.ResponseCreateVariant
// @Failure      500  {object}  dto.ResponseCreateVariant
// @Security     BearerAuth
// @Router       /admin/variants/{id} [patch]
func (v *VariantHandler) UpdateVariant(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	var req dto.VariantRequest

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseUpdateVariant{
			Success: false,
			Message: "invalid id",
			Error:   err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseUpdateVariant{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	data, err := v.service.UpdateVariant(ctx.Request.Context(), req, id)
	if err != nil {
		if err.Error() == "variant not found" {
			ctx.JSON(http.StatusNotFound, dto.ResponseUpdateVariant{
				Success: false,
				Message: "not found",
				Error:   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, dto.ResponseUpdateVariant{
				Success: false,
				Message: "internal server error",
				Error:   err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseUpdateVariant{
		Success: true,
		Message: "success update variant",
		Result:  data,
	})
}

// Delete Variant By Id godoc
// @Summary      Delete variant by id
// @Description  Delete variant by id
// @Tags         Variants
// @Produce      json
// @Param        id   path  int  true  "variant id"
// @Success      200  {object}  dto.ResponseVariant
// @Failure      400  {object}  dto.ResponseVariant
// @Failure      404  {object}  dto.ResponseVariant
// @Failure      500  {object}  dto.ResponseVariant
// @Security     BearerAuth
// @Router       /admin/variants/{id} [delete]
func (v *VariantHandler) DeleteVariant(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseVariant{
			Success: false,
			Message: "invalid id",
			Error:   err.Error(),
		})
		return
	}

	if err := v.service.DeleteVariant(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseVariant{
			Success: false,
			Message: "failed to delete variant",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseVariant{
		Success: true,
		Message: "success delete variant",
	})
}
