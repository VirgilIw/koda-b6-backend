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

// Create New size godoc
// @Summary      Create New size
// @Description  Create New size
// @Tags         Sizes
// @Accept       json
// @Produce      json
// @Param req body dto.SizeRequest true "Create New Size"
// @Success      200  {object}  dto.ResponseSizeCreate
// @Failure      500  {object}  dto.ResponseSizeCreate
// @Router       /admin/sizes [post]
func (h *SizesHandler) CreateSize(ctx *gin.Context) {
	var req dto.SizeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseSizeCreate{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	size, err := h.service.CreateSize(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseSizeCreate{
			Success: false,
			Message: "failed to create size",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseSizeCreate{
		Success: true,
		Message: "success create new size",
		Result:  size,
	})
}

// Get all sizes godoc
// @Summary      get all sizes
// @Description  Get all sizes
// @Tags         Sizes
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.ResponseSizes
// @Failure      500  {object}  dto.ResponseSizes
// @Router       /admin/sizes [get]
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
// @Router       /admin/sizes/{id} [get]
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

// Update size by ID godoc
// @Summary      update size
// @Description  update size by id
// @Tags         Sizes
// @Accept       json
// @Produce      json
// @Param        id        path int            true "size id"
// @Param        request   body dto.SizeUpdateRequest true "update size request"
// @Success      200       {object} dto.ResponseSize
// @Failure      400       {object} dto.ResponseSize
// @Failure      404       {object} dto.ResponseSize
// @Failure      500       {object} dto.ResponseSize
// @Router       /admin/sizes/{id} [patch]
func (h *SizesHandler) UpdateSize(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, dto.ResponseSize{
			Success: false,
			Message: "invalid size id",
		})
		return
	}

	var req dto.SizeUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseSize{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	updatedSize, err := h.service.UpdateSize(ctx.Request.Context(), id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseSize{
			Success: false,
			Message: "failed to update size",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseSize{
		Success: true,
		Message: "success update size",
		Result:  updatedSize,
	})
}

// Delete Size godoc
// @Summary      Delete size
// @Description  Delete size
// @Tags         Sizes
// @Accept       json
// @Produce      json
// @Param        id    path 	int true "size id"
// @Success      200  {object}  dto.ResponseSize
// @Failure      500  {object}  dto.ResponseSize
// @Router       /admin/sizes/{id} [delete]
func (h *SizesHandler) DeleteSizeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseSize{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	if err = h.service.DeleteSizeById(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseSize{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseSize{
		Success: true,
		Message: "success deleted size",
	})
}
