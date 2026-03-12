package handler

import (
	"net/http"

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
