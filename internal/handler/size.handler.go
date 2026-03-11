package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type SizesHandler struct {
	service *service.SizesService
}

func NewSizesHandler(service *service.SizesService) *SizesHandler {
	return &SizesHandler{
		service: service,
	}
}

// Get all sizes godoc
// @Summary      get all sizes
// @Description  Get all sizes
// @Tags         Sizes
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.ResponseSizes
// @Failure      500  {object}  dto.ResponseSizes
// @Router       /sizes [get]
func (h *SizesHandler) GetSizes(ctx *gin.Context) {
	sizes, err := h.service.GetSizes(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseSizes{
			Success: false,
			Message: "failed to get sizes",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseSizes{
		Success: true,
		Message: "success get all size",
		Result:  sizes,
	})
}

// Get size by ID godoc
// @Summary      get size by id
// @Description  retrieve size by id
// @Tags         Sizes
// @Accept       json
// @Produce      json
// @Param        id   path int true "size id"
// @Success      200  {object}  dto.ResponseSize
// @Failure      400  {object}  dto.ResponseSize
// @Failure      404  {object}  dto.ResponseSize
// @Failure      500  {object}  dto.ResponseSize
// @Router       /sizes/{id} [get]
func (h *SizesHandler) GetSizeByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseSize{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	size, err := h.service.GetSizeByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseSize{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseSize{
		Success: true,
		Message: "success get size",
		Result:  size,
	})
}
